package ui

type HTTPServer struct{}

func NewHTTP() *HTTPServer {
	return &HTTPServer{}
}

func (HTTPServer) UseService(s Service) {

}

func (HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
