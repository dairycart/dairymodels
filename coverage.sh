# delete previous coverage report, if it exists
if [ -f coverage.out ]
then
    rm coverage.out
fi

# run tests
go test github.com/dairycart/dairymodels/v1 -coverprofile=coverage.out
go tool cover -html=coverage.out

# delete the new coverage report, if it exists, so I don't accidentally commit it to the repo somehow
if [ -f coverage.out ]
then
    rm coverage.out
fi