package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabahealthnrceporderquery 订单详情查询接口
// alibaba.health.nr.cep.order.query
//
// 订单详情查询接口
func Alibabahealthnrceporderquery(clt *core.SDKClient, req *alihealth2.AlibabahealthnrceporderqueryAPIRequest, session string) (*alihealth2.AlibabahealthnrceporderqueryAPIResponse, error) {
	var resp alihealth2.AlibabahealthnrceporderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
