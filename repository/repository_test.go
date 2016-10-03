package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/hotolab/exago-svc/mocks"
)

var (
	repo = "github.com/foo/bar"
	data = `{"codestats":{"Assertion":0,"CLOC":48,"ExportedFunction":9,"ExportedMethod":6,"Function":15,"FunctionLOC":197,"GoStatement":0,"IfStatement":28,"Import":19,"Interface":2,"LOC":285,"Method":6,"MethodLOC":6,"NCLOC":237,"NOF":3,"Struct":3,"SwitchStatement":1,"Test":1},"projectrunner":{"coverage":{"label":"Code Coverage","data":{"packages":{"codename":{"name":"codename","path":"github.com/jgautheron/codename-generator","coverage":42.857142857142854},"codename_test":{"name":"codename_test","path":"","coverage":0}},"coverage":21.428571428571427},"raw_output":"","execution_time":0,"error":null},"download":{"label":"Go Get","data":null,"raw_output":"","execution_time":13011,"error":null},"goprove":{"label":"Go Prove","data":{"failed":[{"category":"extraCredit","desc":"Blackbox Tests: In addition to standard tests, does the project have blackbox tests?","name":"hasBlackboxTests"},{"category":"extraCredit","desc":"Benchmarks: In addition to tests, does the project have benchmarks?","name":"hasBenches"}],"passed":[{"category":"minimumCriteria","desc":"README Presence: Does the project's include a documentation entrypoint?","name":"hasReadme"},{"category":"minimumCriteria","desc":"Licensed: Does the project have a license?","name":"hasLicense"},{"category":"minimumCriteria","desc":"golint Correctness: Is the linter satisfied?","name":"isLinted"},{"category":"minimumCriteria","desc":"gofmt Correctness: Is the code formatted correctly?","name":"isFormatted"},{"category":"minimumCriteria","desc":"Directory Names and Packages Match: Does each package \u003cpkg\u003e statement's package name match the containing directory name?","name":"isDirMatch"},{"category":"goodCitizen","desc":"Contribution Process: Does the project document a contribution process?","name":"hasContributing"},{"category":"minimumCriteria","desc":"go tool vet Correctness: Is the Go vet satisfied?","name":"isVetted"}]},"raw_output":"","execution_time":2.61000591e+08,"error":null},"test":{"label":"Go Test","data":[{"name":"github.com/jgautheron/codename-generator","execution_time":0.023,"success":true,"tests":[{"name":"TestCodenameGeneration","execution_time":0.02,"passed":true}]}],"raw_output":"=== RUN   TestCodenameGeneration\n--- PASS: TestCodenameGeneration (0.02s)\nPASS\nok  \tgithub.com/jgautheron/codename-generator\t0.023s\n","execution_time":8.19756506e+08,"error":null},"thirdparties":{"label":"Go List (finds all 3rd parties)","data":null,"raw_output":"","execution_time":1.17644683e+08,"error":null}},"lintmessages":{"codename.go":{"golint":[{"col":2,"line":29,"message":"exported const SuperbFilePath should have comment (or a comment on this block) or be unexported","severity":"warning"},{"col":6,"line":33,"message":"exported type FormatType should have comment or be unexported","severity":"warning"},{"col":6,"line":34,"message":"exported type JSONData should have comment or be unexported","severity":"warning"},{"col":1,"line":36,"message":"comment on exported function Get should be of the form \"Get ...\"","severity":"warning"}]},"words.go":{"dupl":[{"col":0,"line":115,"message":"duplicate of words.go:141-151","severity":"warning"},{"col":0,"line":141,"message":"duplicate of words.go:115-125","severity":"warning"},{"col":0,"line":81,"message":"duplicate of words.go:101-110","severity":"warning"},{"col":0,"line":101,"message":"duplicate of words.go:81-90","severity":"warning"}],"gofmt":[{"col":0,"line":1,"message":"file is not gofmted with -s","severity":"warning"}],"golint":[{"col":5,"line":72,"message":"var _dataSuperbJson should be _dataSuperbJSON","severity":"warning"},{"col":6,"line":74,"message":"func dataSuperbJsonBytes should be dataSuperbJSONBytes","severity":"warning"},{"col":6,"line":81,"message":"func dataSuperbJson should be dataSuperbJSON","severity":"warning"},{"col":5,"line":92,"message":"var _dataSuperheroesJson should be _dataSuperheroesJSON","severity":"warning"},{"col":6,"line":94,"message":"func dataSuperheroesJsonBytes should be dataSuperheroesJSONBytes","severity":"warning"},{"col":6,"line":101,"message":"func dataSuperheroesJson should be dataSuperheroesJSON","severity":"warning"}],"gosimple":[{"col":2,"line":234,"message":"'if err != nil { return err }; return nil' can be simplified to 'return err'","severity":"warning"}]}},"metadata":{"image":"https://avatars.githubusercontent.com/u/683888?v=3","description":"A codename generator meant for naming software releases.","stars":13,"last_push":"2015-08-29T20:32:12Z"},"score":{"value":66.15715932988775,"rank":"D","details":[{"name":"thirdparties","score":100,"weight":1.5,"desc":"counts the number of third party libraries","msg":"0 third-party package(s)","url":"https://github.com/jgautheron/gogetimports"},{"name":"codestats","score":69.91451571225885,"weight":1,"desc":"counts lines of code, comments, functions, structs, imports etc in Go code","msg":"48 comments for 285 lines of code","url":"https://github.com/jgautheron/golocc"},{"name":"testcoverage","score":7.179714674571952,"weight":2.45484486000851,"desc":"measures pourcentage of code covered by tests","msg":"coverage is greater or equal to 21.43","url":"https://golang.org/pkg/testing/"},{"name":"testduration","score":99.93434921544139,"weight":1.2,"desc":"measures time taken for testing","msg":"tests took 0.02s","url":"https://golang.org/pkg/testing/"},{"name":"checklist","score":94.56521739130436,"weight":1.8,"desc":"inspects project for the best practices listed in the Go CheckList","msg":"","url":"https://github.com/karolgorecki/goprove","details":[{"name":"projectBuilds","score":100,"weight":1.5,"msg":"check succeeded","url":"https://github.com/matttproud/gochecklist/blob/master/publication/compilation.md"},{"name":"isFormatted","score":100,"weight":3,"msg":"check succeeded","url":"https://github.com/matttproud/gochecklist/blob/master/publication/code_correctness.md"},{"name":"hasReadme","score":100,"weight":3,"msg":"check succeeded","url":"https://github.com/matttproud/gochecklist/blob/master/publication/documentation_entrypoint.md"},{"name":"isDirMatch","score":100,"weight":0.7,"msg":"check succeeded","url":"https://github.com/matttproud/gochecklist/blob/master/publication/dir_pkg_match.md"},{"name":"isLinted","score":100,"weight":0.5,"msg":"check succeeded","url":"https://github.com/matttproud/gochecklist/blob/master/publication/code_correctness.md"},{"name":"hasBenches","score":0,"weight":0.5,"msg":"check failed"}]},{"name":"lintmessages","score":65.45302288378002,"weight":2,"desc":"runs a whole bunch of Go linters","msg":"","url":"https://github.com/alecthomas/gometalinter","details":[{"name":"gosimple","score":90.00876262522593,"weight":1.5,"desc":"examines Go code and reports constructs that can be simplified","msg":"exceeds the warnings/LOC threshold","url":"https://github.com/dominikh/go-simple"},{"name":"gofmt","score":53.17515301305708,"weight":3,"desc":"detects if Go code is incorrectly formatted","msg":"exceeds the warnings/LOC threshold","url":"https://golang.org/cmd/gofmt/"}]}]},"execution_time":"3s","last_update":"2016-10-03T22:12:37.780273283+02:00"}`
)

func TestIsNotLoaded(t *testing.T) {
	rp := &Repository{
		Name: repo,
	}
	loaded := rp.IsLoaded()
	if loaded {
		t.Errorf("The repository %s should not be loaded", rp.Name)
	}
}

func TestIsLoaded(t *testing.T) {
	rp, err := loadStubRepo()
	if err != nil {
		t.Error(err)
	}

	loaded := rp.IsLoaded()
	if !loaded {
		t.Errorf("The repository %s should be loaded", rp.Name)
	}
}

func TestStartTimeSet(t *testing.T) {
	rp := &Repository{
		Name: repo,
	}
	now := time.Now()
	rp.SetStartTime(now)
	if now != rp.startTime {
		t.Error("Got the wrong time")
	}
}

func loadStubRepo() (*Repository, error) {
	rhMock := mocks.RepositoryHost{}
	dbMock := mocks.Database{}
	dbMock.On("Get",
		[]byte(fmt.Sprintf("%s-%s", repo, "")),
	).Return([]byte(data), nil)

	rp := &Repository{
		Name:           repo,
		DB:             dbMock,
		RepositoryHost: rhMock,
	}
	if err := rp.Load(); err != nil {
		return nil, fmt.Errorf("Got error while loading data: %v", err)
	}
	return rp, nil
}
