package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	redisAddr := os.Getenv("REDIS_ADDR")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	defer redisClient.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		status := map[string]string{
			"postgres": "connected",
			"redis":    "connected",
		}

		if err := db.PingContext(ctx); err != nil {
			status["postgres"] = "unavailable"
		}

		if err := redisClient.Ping(ctx).Err(); err != nil {
			status["redis"] = "unavailable"
		}

		w.Header().Set("Content-Type", "application/json")

		if status["postgres"] != "connected" ||
			status["redis"] != "connected" {
			w.WriteHeader(http.StatusServiceUnavailable)
		}

		json.NewEncoder(w).Encode(status)
	})

	log.Println("Go API running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
