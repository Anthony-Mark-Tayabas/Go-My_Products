package middleware

import (
	"database/sql"
	"encoding/json"
	"my-products/models"
	"log"
	"os"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	err := godotenv.Load(".env")
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	db, err = sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS products (
		name TEXT PRIMARY KEY,
		type TEXT,
		picture TEXT,
		price REAL,
		description TEXT
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO products (name, type, picture, price, description) VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(query, product.Name, product.Type, product.PictureURL, product.Price, product.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	query := `SELECT name, type, picture, price, description FROM products`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Name, &product.Type, &product.PictureURL, &product.Price, &product.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var product models.Product
	query := `SELECT name, type, picture, price, description FROM products WHERE name=$1`
	err := db.QueryRow(query, name).Scan(&product.Name, &product.Type, &product.PictureURL, &product.Price, &product.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE products SET type=$1, picture=$2, price=$3, description=$4 WHERE name=$5`
	_, err = db.Exec(query, product.Type, product.PictureURL, product.Price, product.Description, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	query := `DELETE FROM products WHERE name=$1`
	_, err := db.Exec(query, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}