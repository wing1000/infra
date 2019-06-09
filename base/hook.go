package base

import (
	log "github.com/sirupsen/logrus"
	"imooc.com/resk/infra"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

var callbacks []func()

func Register(fn func()) {
	callbacks = append(callbacks, fn)
}

type HookStarter struct {
	infra.BaseStarter
}

func (s *HookStarter) Init(ctx infra.StarterContext) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for {
			c := <-sigs
			log.Info("notify: ", c)
			for _, fn := range callbacks {
				fn()
			}
			break
			os.Exit(0)
		}
	}()

}

func (s *HookStarter) Start(ctx infra.StarterContext) {
	starters := infra.GetStarters()

	for _, s := range starters {
		typ := reflect.TypeOf(s)
		log.Infof("【Register Notify Stop】:%s.Stop()", typ.String())
		Register(func() {
			s.Stop(ctx)
		})
	}

}
