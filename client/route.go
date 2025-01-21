package main

import (
	"encoding/json"
	"io"
	"net/http"

	pb "github.com/ronexlemon/godocker/common/api"
)

// start mux serve
type clientHandler struct {
	//mux
	client pb.StoreServiceClient
}

func NewClientHandler(client pb.StoreServiceClient) *clientHandler {
	return &clientHandler{client}

}

func (c *clientHandler) registerRoute(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/menu", c.GetMenuOrders)

}
func (c *clientHandler) GetMenuOrders(w http.ResponseWriter, r *http.Request) {
	// Call the gRPC GetMenu method
	stream, err := c.client.GetMenu(r.Context(), &pb.MenuRequest{})
	if err != nil {
		http.Error(w, "Failed to get menu: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var allOrders []*pb.OrderMenu

	// Loop to receive the stream from gRPC
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, "Error while receiving stream: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Append the received order to the list
		allOrders = append(allOrders, msg)
	}

	// Serialize the orders to JSON and write to the response
	if err := json.NewEncoder(w).Encode(allOrders); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
