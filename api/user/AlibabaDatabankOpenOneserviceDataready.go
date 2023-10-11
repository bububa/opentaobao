package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AlibabaDatabankOpenOneserviceDataready 瓴羊DaaS消费者增长CGP查询DataReady
// alibaba.databank.open.oneservice.dataready
//
// 瓴羊DaaS消费者增长CGP取数接口
func AlibabaDatabankOpenOneserviceDataready(clt *core.SDKClient, req *user.AlibabaDatabankOpenOneserviceDatareadyAPIRequest, session string) (*user.AlibabaDatabankOpenOneserviceDatareadyAPIResponse, error) {
	var resp user.AlibabaDatabankOpenOneserviceDatareadyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
