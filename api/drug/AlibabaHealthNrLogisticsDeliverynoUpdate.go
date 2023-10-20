package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// Alibabahealthnrlogisticsdeliverynoupdate 上传订单同城快递单号
// alibaba.health.nr.logistics.deliveryno.update
//
// 上传订单同城快递单号
func Alibabahealthnrlogisticsdeliverynoupdate(clt *core.SDKClient, req *drug.AlibabahealthnrlogisticsdeliverynoupdateAPIRequest, session string) (*drug.AlibabahealthnrlogisticsdeliverynoupdateAPIResponse, error) {
	var resp drug.AlibabahealthnrlogisticsdeliverynoupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
