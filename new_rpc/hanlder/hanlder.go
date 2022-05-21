package hanlder

const HelloServiceName = "handler/HelloService"

type NewHelloService struct {
}

func (h *NewHelloService) Hello(req string, reply *string) error {
	*reply = "hello " + req
	return nil

}
