package engine



type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan chan Item
}



type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	//ConfigureMasterWorkerChan(chan Request)
	WorkerChan() chan Request

	Run()
}

type ReadyNotifier interface {
	WorkReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//in := make(chan Request)
	out := make(chan ParserResult)

	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in, out)
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {

			go func() {
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var visitedUrls = make(map[string]bool)
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true

	return false
}


func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

//func createWorker(in chan Request, out chan ParserResult) {
//	go func() {
//		for {
//			request := <-in
//			result, err := worker(request)
//			if err != nil {
//				continue
//			}
//			out <- result
//		}
//	}()
//}
