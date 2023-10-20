package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// Alibabahealthnrlogisticswaybillget 电子面单查询接口
// alibaba.health.nr.logistics.waybill.get
//
// 商家登录后根据订单号查询物流单号及电子面单信息
func Alibabahealthnrlogisticswaybillget(clt *core.SDKClient, req *drug.AlibabahealthnrlogisticswaybillgetAPIRequest, session string) (*drug.AlibabahealthnrlogisticswaybillgetAPIResponse, error) {
	var resp drug.AlibabahealthnrlogisticswaybillgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
