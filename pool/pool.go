package pool

import (
	"time"

	exago "github.com/hotolab/exago-svc"
	. "github.com/hotolab/exago-svc/config"
	"github.com/hotolab/exago-svc/repository/model"
	"github.com/jeffail/tunny"
)

const (
	SendTimeout = time.Second * 280
)

var (
	// Make sure it satisfies the interface.
	_ model.Pool = (*PoolRunner)(nil)
)

type PoolRunner struct {
	pool   *tunny.WorkPool
	config exago.Config
}

func New(options ...exago.Option) (model.Pool, error) {
	var p PoolRunner
	for _, option := range options {
		option.Apply(&p.config)
	}
	pool, err := tunny.CreatePool(Config.PoolSize, p.config.RepositoryProcessor).Open()
	if err != nil {
		return nil, err
	}
	p.pool = pool
	return &p, nil
}

func (pr *PoolRunner) PushSync(repo, branch, goversion string) (model.Record, error) {
	value, _ := pr.pool.SendWork([]string{repo, branch, goversion})
	switch value.(type) {
	case error:
		return nil, value.(error)
	default:
		return value.(model.Record), nil
	}
	return nil, nil
}

func (pr *PoolRunner) PushAsync(repo, branch, goversion string) {
	pr.pool.SendWorkAsync([]string{repo, branch, goversion}, nil)
}

func (pr *PoolRunner) WaitUntilEmpty() {
	for {
		time.Sleep(1 * time.Second)
		if pr.pool.NumPendingAsyncJobs() == 0 {
			return
		}
	}
}

func (pr *PoolRunner) Stop() {
	pr.pool.Close()
}