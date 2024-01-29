env "local" {
  src = "file://sql/schema.sql"
  url = "postgres://short:short@localhost:5432/short?sslmode=disable"
  migration {
    dir = "file://sql/migrations"
  }
}
