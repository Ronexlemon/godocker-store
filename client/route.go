package main

import (
	"encoding/json"
	"io"
	"net/http"

	pb "github.com/ronexlemon/godocker/common/api"
)


type clientHandler struct {
	
	client pb.StoreServiceClient
}

func NewClientHandler(client pb.StoreServiceClient) *clientHandler {
	return &clientHandler{client}

}

func (c *clientHandler) registerRoute(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/menu", c.GetMenuOrders)
	mux.HandleFunc("POST /api/order",c.placeOrder)

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


//function placeOrder


func (c *clientHandler) placeOrder(w http.ResponseWriter, r *http.Request){
	// Parse the JSON request body
	var order pb.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Failed to parse order: "+err.Error(), http.StatusBadRequest)
		return
		}
		// Call the gRPC PlaceOrder method
		receipt, err := c.client.PlaceOrder(r.Context(), &pb.Order{Id: order.Id,Name: order.Name})
		if err !=nil{
			http.Error(w,"Failed to place Order"+err.Error(),http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(receipt)
}