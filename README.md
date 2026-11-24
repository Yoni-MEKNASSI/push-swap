# push-swap

Projet Go (bibliothèque standard uniquement).

    go build -o push-swap ./cmd/push-swap
    go build -o checker ./cmd/checker
    ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
    go test ./...

- `internal/stack` : les piles a et b, les 11 instructions, la lecture des arguments
- `cmd/push-swap` : l'algorithme de tri (cas à la main pour ≤ 5, radix au-delà)
- `cmd/checker` : exécute les instructions et affiche OK ou KO
