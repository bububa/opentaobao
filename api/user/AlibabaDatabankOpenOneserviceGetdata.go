package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AlibabaDatabankOpenOneserviceGetdata 瓴羊DaaS消费者运营CGP取数接口
// alibaba.databank.open.oneservice.getdata
//
// 瓴羊DaaS消费者运营CGP取数接口
func AlibabaDatabankOpenOneserviceGetdata(clt *core.SDKClient, req *user.AlibabaDatabankOpenOneserviceGetdataAPIRequest, session string) (*user.AlibabaDatabankOpenOneserviceGetdataAPIResponse, error) {
	var resp user.AlibabaDatabankOpenOneserviceGetdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
