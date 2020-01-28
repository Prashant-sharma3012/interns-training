package server

import "net/http"

func (s *Server) RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("its in progress"))
}
