module imooc.com/resk

go 1.12

//无法访问的原因，替换golang.org源为github.com源
replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.37.2
	golang.org/x/build => github.com/golang/build v0.0.0-20190327004547-5a2224f3eb52
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190321205749-f0864edee7f3
	golang.org/x/image => github.com/golang/image v0.0.0-20190321063152-3fc05d484e9f
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190319155245-9487ef54b94a
	golang.org/x/net => github.com/golang/net v0.0.0-20190327025741-74e053c68e29
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190319182350-c85d3e98c914
	golang.org/x/perf => github.com/golang/perf v0.0.0-20190312170614-0655857e383f
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190322080309-f49334f85ddc
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190327011446-79af862e6737
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.3.0
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190321212433-e79c0c59cdb5
	google.golang.org/grpc => github.com/grpc/grpc-go v1.19.1

)

require (
	github.com/go-playground/locales v0.12.1
	github.com/go-playground/universal-translator v0.16.0
	github.com/go-redsync/redsync v1.2.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/json-iterator/go v1.1.6
	github.com/kataras/iris v11.1.1+incompatible
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/mattn/go-colorable v0.1.1
	github.com/prometheus/common v0.2.0
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/segmentio/ksuid v1.0.2
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/sirupsen/logrus v1.4.1
	github.com/smartystreets/goconvey v0.0.0-20190306220146-200a235640ff
	github.com/tietang/dbx v0.0.0-20190314092248-c144fca0b077
	github.com/tietang/go-eureka-client/eureka v0.0.0-20190327071554-ed5a2bb78851
	github.com/tietang/go-utils v0.1.2
	github.com/tietang/props v2.3.0+incompatible
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	gopkg.in/go-playground/validator.v9 v9.27.0
)
