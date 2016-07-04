package lambda

import (
	"encoding/json"

	"github.com/exago/svc/repository/model"
)

var codeStatsCmd = &cmd{
	name:      model.CodeStatsName,
	unMarshal: unMarshalCodeStats,
}

func (l Runner) FetchCodeStats() (interface{}, error) {
	codeStatsCmd.ctxt = context{
		Repository: l.repository,
		Cleanup:    l.shouldCleanup,
	}
	return codeStatsCmd.Data()
}

func unMarshalCodeStats(l *cmd, b []byte) (interface{}, error) {
	var cs model.CodeStats
	err := json.Unmarshal(b, &cs)
	return cs, err
}