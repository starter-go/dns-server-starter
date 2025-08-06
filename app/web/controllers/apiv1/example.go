package apiv1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/dns-server-starter/app/data/dao"
	"github.com/starter-go/dns-server-starter/app/data/dxo"
	"github.com/starter-go/dns-server-starter/app/web/dto"
	"github.com/starter-go/dns-server-starter/app/web/vo"
	"github.com/starter-go/libgin"
)

// ExampleController ...
type ExampleController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")
	Dao    dao.ExampleDAO   //starter:inject("#")
}

func (inst *ExampleController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *ExampleController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *ExampleController) route(rp libgin.RouterProxy) error {

	rp = rp.For("example")

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handle)
	rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *ExampleController) handle(c *gin.Context) {
	req := &myExampleRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *ExampleController) handleGetOne(c *gin.Context) {
	req := &myExampleRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

////////////////////////////////////////////////////////////////////////////////

type myExampleRequest struct {
	context    *gin.Context
	controller *ExampleController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ExampleID
	body1 vo.Example
	body2 vo.Example
}

func (inst *myExampleRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.ExampleID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myExampleRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myExampleRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myExampleRequest) doNOP() error {
	return nil
}

func (inst *myExampleRequest) doGetOne() error {
	id := inst.id
	o1, err := inst.controller.Dao.Find(nil, id)
	if err != nil {
		return err
	}
	o2 := &dto.Example{
		ID:  o1.ID,
		Foo: o1.Foo,
		Bar: o1.Bar,
	}
	inst.body2.Items = []*dto.Example{o2}
	return nil
}
