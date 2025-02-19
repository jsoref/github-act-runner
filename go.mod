module main

go 1.13

require (
	github.com/AlecAivazis/survey/v2 v2.2.7
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.2.0
	github.com/mtibben/androiddnsfix v0.0.0-20200907095054-ff0280446354
	github.com/nektos/act v0.2.22
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.1.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/nektos/act => ./act
