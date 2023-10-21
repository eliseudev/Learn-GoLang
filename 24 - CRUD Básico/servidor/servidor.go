package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if err = json.Unmarshal(bodyRequest, &usuario); err != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into usuarios(nome, email) values(?, ?)")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return

	}

	idInsert, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o ID inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(([]byte(fmt.Sprintf("Usuãrio inserido com sucesso! ID: %d", idInsert))))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar copm o banco!"))
		return
	}
	defer db.Close()

	lines, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}
	defer lines.Close()

	var usuarios []usuario
	for lines.Next() {
		var usuario usuario
		if err := lines.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}

	line, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}

	var usuario usuario
	if line.Next() {
		if err := line.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("Erro ao cconverter usuário para JSON!"))
		return
	}
}
