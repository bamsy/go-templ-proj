package handler

import (
	"net/http"

	"github.com/bamsy23/go-templ-proj/model"
	"github.com/bamsy23/go-templ-proj/view/user"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(w http.ResponseWriter, r *http.Request) {
	u, _ := r.Context().Value("user").(model.User)
	user.Show(u).Render(r.Context(), w)
}
