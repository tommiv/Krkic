package training

import (
    "sort"
    "krkic/model"
)

type ByFirstAttempt []model.Bojan

func (input ByFirstAttempt) Len() int {
    return len(input)
}

func (input ByFirstAttempt) Swap(i, j int) {
    input[i], input[j] = input[j], input[i]
}

func (input ByFirstAttempt) Less(i, j int) bool {
    a := input[i].Attempts[0].BojanedAt
    b := input[j].Attempts[0].BojanedAt
    return a.Before(b)
}

func SortByFirstAttempt(input []model.Bojan) []model.Bojan {
    sort.Sort(ByFirstAttempt(input))
    return input
}
