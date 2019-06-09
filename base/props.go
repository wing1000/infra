package base

import (
	log "github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
	"imooc.com/resk/infra"
	"sync"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	Check(props)
	return props
}

type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	props = ctx.Props()
	log.Info("初始化配置.")
	GetSystemAccount()
}

type SystemAccount struct {
	AccountNo   string
	AccountName string
	UserId      string
	Username    string
}

var systemAccount *SystemAccount
var systemAccountOnce sync.Once

func GetSystemAccount() *SystemAccount {
	systemAccountOnce.Do(func() {
		systemAccount = new(SystemAccount)
		err := kvs.Unmarshal(Props(), systemAccount, "system.account")
		if err != nil {
			panic(err)
		}
	})
	return systemAccount
}

func GetEnvelopeActivityLink() string {
	link := Props().GetDefault("envelope.link", "/v1/envelope/link")
	return link
}

func GetEnvelopeDomain() string {
	domain := Props().GetDefault("envelope.domain", "http://localhost")
	return domain
}
