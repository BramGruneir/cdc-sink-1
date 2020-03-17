package main

import (
    "database/sql"
    "fmt"
    "log"
    "flag"

    _ "github.com/lib/pq"
)

var connectionString = flag.String("conn", "postgresql://root@localhost:26257/defaultdb?sslmode=disable", "cockroach connection string")

func main() {
    db, err := sql.Open("postgres", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")
    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }

    // Create the "accounts" table.
    if _, err := db.Exec(
        "CREATE TABLE IF NOT EXISTS accounts (id INT PRIMARY KEY, balance INT)"); err != nil {
        log.Fatal(err)
    }

    // Insert two rows into the "accounts" table.
    if _, err := db.Exec(
        "INSERT INTO accounts (id, balance) VALUES (1, 1000), (2, 250)"); err != nil {
        log.Fatal(err)
    }

    // Print out the balances.
    rows, err := db.Query("SELECT id, balance FROM accounts")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    fmt.Println("Initial balances:")
    for rows.Next() {
        var id, balance int
        if err := rows.Scan(&id, &balance); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%d %d\n", id, balance)
    }
}
