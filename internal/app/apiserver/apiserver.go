package apiserver

type APIServer struct {
}

func NewAPIServer() *APIServer {
	return &APIServer{}
}

func (s *APIServer) Start() error {
	return nil
}
