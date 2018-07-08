package rpcdemo

import "errors"

// Service,Method

type DemoService struct {
}

type Args struct {
	A, B int
}

//rpc 函数的参数一定要有两个,一个参数,一个结果
func (demoService DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}
