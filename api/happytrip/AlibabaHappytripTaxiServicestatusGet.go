package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxiservicestatusget 供应商服务开通状态
// alibaba.happytrip.taxi.servicestatus.get
//
// 获取服务供应商在每个地区的服务开通状态、支持的车型等
func Alibabahappytriptaxiservicestatusget(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxiservicestatusgetAPIRequest, session string) (*happytrip.AlibabahappytriptaxiservicestatusgetAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxiservicestatusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
