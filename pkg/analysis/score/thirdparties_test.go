package score_test

import (
	"testing"

	exago "github.com/jgautheron/exago/pkg"

	"github.com/jgautheron/exago/pkg/analysis/score"
)

func TestThirdParties(t *testing.T) {
	var tests = []struct {
		thirdParties []string
		loc          int
		operator     string
		expected     float64
		desc         string
	}{
		{getThirdParties(2), 200, ">", 70, "Two third parties for a small project is pretty common"},
		{getThirdParties(8), 5000, ">", 50, "8 third parties for 5000 LOC is not that bad"},
		{getThirdParties(9), 2000, "<", 50, "For 2k LOC this is proportionally too much"},
		{getThirdParties(10), 400, "<", 30, "Way too much third parties for 400 LOC"},
		{[]string{}, 100, "=", 100, "No third party, then obviously we get the maximum score"},
	}

	for _, tt := range tests {
		d := exago.Data{}
		d.Results = exago.Results{}
		d.Results.ThirdParties.Data = tt.thirdParties
		d.Results.CodeStats.Data = map[string]int{"loc": tt.loc}
		res := score.ThirdPartiesEvaluator().Calculate(d)

		switch tt.operator {
		case "<":
			if res.Score > tt.expected {
				t.Error("Wrong score")
			}
		case ">":
			if res.Score < tt.expected {
				t.Error("Wrong score")
			}
		case "=":
			if res.Score != tt.expected {
				t.Error("Wrong score")
			}
		}
	}
}

func getThirdParties(count int) (tp []string) {
	for i := 0; i < count; i++ {
		tp = append(tp, string(i))
	}
	return tp
}
