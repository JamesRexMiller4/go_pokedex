module github.com/jamesrexmiller4/go_pokedex

go 1.22.1

require internal/apis v1.0.0

replace internal/apis => ./internal/apis

require internal/types v1.0.0

replace internal.apis => ./internal/types
