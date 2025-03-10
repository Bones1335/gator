# Blog Aggregator Guided Project

To use gator, you will need PostgreSQL and GO installed on your machine.

The program can be installed by running `go install https://github.com/Bones1335/gator`.

Make sure to create a config file `.gatorconfig.json` in your home directory on your machine with the following data:
```JSON
{
    "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```

