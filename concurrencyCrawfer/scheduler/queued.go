package scheduler

import "spiders/concurrencyCrawfer/engine"

type QueuedScheduler struct {
	// request
	requestChan chan engine.Request
	// work  把100个channel放到一起
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

// WorkReady 有一个work ready ok,就可以负责去接收request
func (s *QueuedScheduler) WorkReady(w chan engine.Request) {
	s.workerChan <- w
}

// 总控的goroutine
func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		//声明两个队列来保存request
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-s.requestChan: // 将request加到队列
				// 收到request让其排队
				requestQ = append(requestQ, r)
			case w := <-s.workerChan: // 将worker加到队列
				// 收到一个worker让其排队
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest: // 同时又request和worker,把request发给worker
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}

		}
	}()
}
