package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"myapp/config"
	"myapp/models"

	"github.com/gorilla/mux"
)

func GetAllBarang(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(context.Background(), "SELECT id, nama, harga FROM barang")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var barangs []models.Barang
	for rows.Next() {
		var barang models.Barang
		if err := rows.Scan(&barang.ID, &barang.Nama, &barang.Harga); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		barangs = append(barangs, barang)
	}

	json.NewEncoder(w).Encode(barangs)
}

func GetBarangByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var barang models.Barang
	err = config.DB.QueryRow(context.Background(), "SELECT id, nama, harga FROM barang WHERE id=$1", id).Scan(&barang.ID, &barang.Nama, &barang.Harga)
	if err != nil {
		http.Error(w, "Barang not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(barang)
}

func CreateBarang(w http.ResponseWriter, r *http.Request) {
	var barang models.Barang
	if err := json.NewDecoder(r.Body).Decode(&barang); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := config.DB.QueryRow(context.Background(), "INSERT INTO barang (nama, harga) VALUES ($1, $2) RETURNING id", barang.Nama, barang.Harga).Scan(&barang.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(barang)
}

func UpdateBarang(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var barang models.Barang
	if err := json.NewDecoder(r.Body).Decode(&barang); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec(context.Background(), "UPDATE barang SET nama=$1, harga=$2 WHERE id=$3", barang.Nama, barang.Harga, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteBarang(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec(context.Background(), "DELETE FROM barang WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
