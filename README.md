# Push-Swap

Projet d'algorithmique développé en Go (utilisant uniquement la bibliothèque standard). L'objectif est de trier une liste d'entiers sur une pile en utilisant un ensemble restreint d'instructions et un nombre minimal de mouvements.

## Structure du dépôt

Le projet est organisé de la manière suivante :

* `go.mod` : Configuration du module Go (version 1.21).
* `cmd/push-swap/` : Programme principal qui calcule et affiche la séquence d'instructions pour trier les nombres.
* `cmd/checker/` : Programme de vérification qui lit des instructions sur l'entrée standard et valide si la pile est correctement triée (`OK` ou `KO`)[cite: 1].
* `internal/stack/` : Paquet interne contenant la logique des piles (A et B), les 11 instructions et l'analyse des arguments[cite: 1].

## Les 11 instructions disponibles

Le sommet de chaque pile correspond au premier élément (index `0`). Les opérations supportées sont :

* **`sa` / `sb` / `ss`** : Échange les deux premiers éléments du sommet de la pile A, de la pile B, ou des deux en même temps.
* **`pa` / `pb`** : Prend le premier élément du sommet de B (resp. A) et le place au sommet de A (resp. B).
* **`ra` / `rb` / `rr`** : Décale tous les éléments d'une pile vers le haut (le premier élément devient le dernier).
* **`rra` / `rrb` / `rrr`** : Décale tous les éléments d'une pile vers le bas (le dernier élément devient le premier)[cite: 1].

## Fonctionnement de l'algorithme (`push-swap`)

1. **Normalisation (Rangs) :** Les nombres saisis sont convertis en rangs (de `0` à `N-1`) après un tri initial. Cela permet de traiter facilement les nombres négatifs et les grands écarts de valeurs[cite: 1].
2. **Petits nombres ($\le 5$) :** 
   * Pour 2 nombres : un simple `sa`[cite: 1].
   * Pour 3 à 5 nombres : les plus petits éléments sont envoyés dans la pile B, les éléments restants dans A sont triés manuellement, puis les éléments de B sont replacés dans A[cite: 1].
3. **Grandes listes ($> 5$) :** Utilisation d'un tri Radix basé sur les bits. Les nombres sont basculés entre les piles A et B bit par bit jusqu'à ce que l'ensemble soit trié[cite: 1].

## Installation et Compilation

Assurez-vous d'avoir Go installé, puis compilez les binaires depuis la racine du projet :

```bash
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker