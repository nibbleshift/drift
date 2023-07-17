package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/nibbleshift/drift"
	"github.com/nibbleshift/drift/ent"
	"github.com/rs/cors"
)

const defaultPort = "8080"

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func runAuth() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	casdoorsdk.InitConfig(
		"http://localhost:8000",
		"99c54cb29ac856f66c64",                     // client id
		"bae92cffeae4905a6b54d60ea7c68179805202b6", // client secret
		cert,    // certificate
		"drift", // organization
		"drift", // application
	)

	r.Get("/api/getUserInfo", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")
		spew.Dump(code, state)
		token, err := casdoorsdk.GetOAuthToken(code, state)
		if err != nil {
			panic(err)
		}

		claims, err := casdoorsdk.ParseJwtToken(token.AccessToken)
		if err != nil {
			panic(err)
		}

		spew.Dump(claims)

		claims.AccessToken = token.AccessToken
		jsonData, err := json.Marshal(claims)

		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, string(jsonData))
	})

	r.Post("/*", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Post called")
		code := r.URL.Query().Get("code")
		spew.Dump(code)
		jsonData, err := json.Marshal(code)

		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, string(jsonData))
	})
	http.ListenAndServe(":9900", r)
}

func main() {
	client := Open("postgresql://postgres:postgres@127.0.0.1/drift")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx := context.Background()
	err := client.Schema.Create(ctx)

	if err != nil {
		log.Println(err)
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
			"http://localhost:8080",
			"http://localhost:8000",
		},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(drift.NewSchema(client))
	//srv := handler.NewDefaultServer(drift.NewExecutableSchema(drift.Config{Resolvers: &drift.Resolver{}}))

	srv.AddTransport(&transport.Websocket{}) // <---- This is the important part!

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	go runAuth()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

var cert string = `-----BEGIN CERTIFICATE-----
MIIE2TCCAsGgAwIBAgIDAeJAMA0GCSqGSIb3DQEBCwUAMCYxDjAMBgNVBAoTBWFk
bWluMRQwEgYDVQQDDAtjZXJ0Xzg0MG1nOTAeFw0yMzA3MTcwNTUxMzRaFw00MzA3
MTcwNTUxMzRaMCYxDjAMBgNVBAoTBWFkbWluMRQwEgYDVQQDDAtjZXJ0Xzg0MG1n
OTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAMjHfaILiaQ3b507/U3V
aDaojYkEHi5/8VEyxODwIkQr7fRFLyKQRZ9K1dFouOPXbyP1UR735A5xBa8C542W
qYkkprennOsAlAnU5/l/GF+dhr7SdQXOYFMO8nyvLU8oMETQU9nawzf59FNBW1XD
8fn1u7X/Lo6w7rT8wRVy+vAD/q3DWlrbMKqX5rmOfmiStEHpx4KqaAcbfBvpDjLZ
C3MwwKkFOexFflgdXplD72A5oRKss7ZNEqlsFrrVz1g+hnm9vMnjDN20TIQgidgH
Q2ehYYNyKO3HJ/TZZABl3WFSPFohxDxYhGtj4arFt5PeIFwIENxKF1y78U5IzmlZ
aEPFleZjXLZp0I6ee8RwSr2+n4ET+sHLbhPCQB9UuJHRiwWToIhXrkLuxODSJ/lz
Q7j8mZj+G1DpaJRAXFCnpZtDKnIc6U4SWHnPN0VHyXK6d4DB5UE1TpXqicJf9CVZ
Yyv+LkyYb50RnPxHsLHLxtPYaZ+Ad7QZ955nkSWbFuPrJf9KyMDv73r8jcvVcLfX
HLCatvC80XT+6lkkYaVYM0fwytLBxhLe1KP0M4jR9dU68yEDPcj6uaKlZsRWlRnm
KB6rrMSJGoA2zcNA2d3oIcbp8x94exZZ4EuskjgGQ3ys3MaHhQoFh5O4SQzReFta
klKTE6h7xuyUpHDvgyNF5yvhAgMBAAGjEDAOMAwGA1UdEwEB/wQCMAAwDQYJKoZI
hvcNAQELBQADggIBALcO9crf2eImAViJwGyYWUlF1otaQ12Zv/Ty77C3NG2su4nK
cVwyBN5/Akkq5g3LRo8gmOy8yINSC6e4Bly5v6sI2fkb9OugYWNPjkJ5QmH1BDM8
+ZWqxMwuxo0TtNFivrgEYQUq7t+aBZoHzs7jsdYOS9XVl0GC6HLTViDYhfZ4Z4eu
4BgcRc5KloXgpev5V1CXWzuom5Sa0ftwi7mNETg7UJXRKkU0DUNpQVKJtgWWYskC
49fHjjYizInQMkgswl9dI7UhG8bZggnEilgIr8vwSlNuGWH7AizDwSm3txEdROar
1PNFQfguT2Ej4L455axlGZnDLV2kyoaNRmxuzBcR0frDlke7d8WFT/bZ3NcvBxLB
Z/X14OzNCuSN+fde/goBc2erGgFUyUOIGILLWS7MnIkZFC6dv5ionwuo89cVgYro
0WNe8aN6b8B0v2jHe3YPWId7YUir3bkb660ObkM34ohrknWJSyt4eOeQp2iW6TUb
7G6mjYr6gzhPayPqx+MAfRd8Veyp0y4nOmYTF+GmBn3XceUcj1iVIlNxoB727z5u
2agJZKtns7SKN/z2qhq3J/NwyhvSJ8fHxCHzwgVO7RSwpB4pxrdovVWdFWJ3aeSg
ToXSBgbyI00bWxSl3F/i/UVmfcrEfeqRR3vNpXEpvBf+ksmHfJrSIG5doNL8
-----END CERTIFICATE-----`
