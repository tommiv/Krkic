package training

import (
    "net/url"
    "krkic/model"
)

func MergeByURL(input []model.Bojan) []model.Bojan {
    merged := make(map[url.URL]model.Bojan, len(input))

    for _, bojan := range input {
        if existing, ok := merged[bojan.URL]; ok {
            attempt := bojan.Attempts[0]
            existing.Attempts = append(existing.Attempts, attempt)
        } else {
            merged[bojan.URL] = bojan
        }
    }

    result := make([]model.Bojan, len(merged))

    i := 0
    for _, val := range merged {
        result[i] = val
        i++
    }

    return result
}
