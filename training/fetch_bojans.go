package training

import (
    "krkic/model"
    "krkic/fetch"
    "github.com/ivpusic/grpool"
    log "gopkg.in/Sirupsen/logrus.v0"
)

func FetchBojans(jobs []model.FetcherJob) []*model.Bojan {
    // TODO: configurable pool size
    pool := grpool.NewPool(12, len(jobs))
    pool.WaitCount(len(jobs))

    var bojans []*model.Bojan;

    for _, item := range jobs {
        job := item // should pin job in the closure

        pool.JobQueue <- func() {
            result, err := fetch.GetHashByUrl(job.URL)
            if err == nil {
                bojans = append(bojans, result)
            } else {
                log.Error(err)
            }
            pool.JobDone()
        }
    }

    pool.WaitAll()
    pool.Release()

    return bojans
}
