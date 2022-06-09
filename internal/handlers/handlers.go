package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/catehulu/rigged-coin/internal/config"
	"github.com/catehulu/rigged-coin/internal/driver"
	"github.com/catehulu/rigged-coin/internal/repository"
	"github.com/catehulu/rigged-coin/internal/repository/dbrepo"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

type jsonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewMongoDBRepo(db.MongoDB, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// PostBoard posts the board pieces
func (m *Repository) PostBoards(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		log.Print("Error parsing form")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("an error occured"))
		return
	}

	if !r.Form.Has("col") || !r.Form.Has("row") {
		w.WriteHeader(http.StatusBadRequest)
		res := jsonResponse{
			Status:  400,
			Message: "All field is required",
			Data:    nil,
		}
		err := responseJson(w, res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Print("Error")
			w.Write([]byte("An error occured"))
			return
		}
		return
	}

	id := r.Form.Get("id")
	col, err := strconv.ParseInt(r.Form.Get("col"), 10, 64)
	if err != nil {
		log.Print("Error parsing form")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("an error occured"))
		return
	}

	row, err := strconv.ParseInt(r.Form.Get("row"), 10, 64)
	if err != nil {
		log.Print("Error parsing form")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("an error occured"))
		return
	}

	board := m.DB.FindBoard(id)
	if board == nil {
		res := jsonResponse{
			Status:  404,
			Message: "No board found",
			Data:    nil,
		}
		err := responseJson(w, res)
		if err != nil {
			log.Print("Error")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An error occured"))
		}
		return
	}

	if col+1 > int64(board.Size) || row+1 > int64(board.Size) {
		w.WriteHeader(http.StatusBadRequest)
		res := jsonResponse{
			Status:  400,
			Message: "Bad request",
			Data:    nil,
		}
		err := responseJson(w, res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Print("Error")
			w.Write([]byte("An error occured"))
			return
		}
		return
	}

	prize := board.GetPiece(int(col), int(row))
	board.CheckPrize()

	if board.ObtainedPrize != -1 {
		res := jsonResponse{
			Status:  200,
			Message: "Board already done",
			Data:    board,
		}
		responseJson(w, res)
	} else {
		res := jsonResponse{
			Status:  200,
			Message: "Success",
			Data:    prize,
		}
		responseJson(w, res)
	}
	m.DB.UpdateBoard(board)
}

// GetBoard gets the board detail
func (m *Repository) GetBoards(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		log.Print("Error parsing form")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("an error occured"))
		return
	}

	id := r.Form.Get("id")
	board := m.DB.FindBoard(id)
	if board == nil {
		res := jsonResponse{
			Status:  404,
			Message: "No board found",
			Data:    nil,
		}
		err := responseJson(w, res)
		if err != nil {
			log.Print("Error")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An error occured"))
			return
		}
	}

	res := jsonResponse{
		Status:  201,
		Message: "Success",
		Data:    board,
	}
	responseJson(w, res)
}

func responseJson(w http.ResponseWriter, jres jsonResponse) error {
	out, err := json.Marshal(jres)
	if err != nil {
		return err
	}
	w.Write(out)
	return nil
}
