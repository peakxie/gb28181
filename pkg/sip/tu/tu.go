package tu

import (
	"context"
	"sync"
	"time"

	"github.com/peakxie/gb28181/pkg/sip/trans"
)

type TU struct {
	ts       sync.Map //tid:*trans.Transactioner
	evChan   chan *trans.Event
	stopCtx  context.Context
	stopFunc context.CancelFunc
}

func NewTU() *TU {
	t := &TU{
		evChan: make(chan *trans.Event, 1000),
	}
	t.stopCtx, t.stopFunc = context.WithCancel(context.Background())
	return t
}

//添加事件
func (t *TU) SetEvent(ev *trans.Event) {
	select {
	case <-t.stopCtx.Done():
	case t.evChan <- ev:
	}
}

func (t *TU) Run() {
	go t.checkTimeoutLoop()
	t.processLoop()
}

//读取event处理
func (t *TU) processLoop() {
	defer t.stopFunc()

	for {
		select {
		case <-t.stopCtx.Done():
			return
		case ev := <-t.evChan:
			var tr trans.Transactioner
			if v, ok := t.ts.Load(ev.Tid); ok {
				if vv, ok := v.(trans.Transactioner); ok {
					tr = vv
				}
			}

			if tr == nil {
				switch ev.Type {
				case trans.RCV_REQINVITE:
					tr = trans.NewIst(ev.Tid, NewIstcb())
				case trans.RCV_REQUEST:
					tr = trans.NewNist(ev.Tid, NewNistcb())
				case trans.SND_REQINVITE:
					tr = trans.NewIct(ev.Tid, NewIctcb())
				case trans.SND_REQUEST:
					tr = trans.NewNict(ev.Tid, NewNictcb())
				default:
					continue
				}
				t.ts.Store(ev.Tid, tr)
			}

			go tr.Do(ev)
		}
	}
}

//循环检测超时
func (t *TU) checkTimeoutLoop() {
	for {
		select {
		case <-t.stopCtx.Done():
			return
		default:
		}

		time.Sleep(time.Second * 60)
		t.ts.Range(func(key, value interface{}) bool {
			tr := value.(trans.Transactioner)
			//清除结束的事务
			if tr.IsTerminated() {
				t.ts.Delete(key)
			}
			return true
		})
	}

}
