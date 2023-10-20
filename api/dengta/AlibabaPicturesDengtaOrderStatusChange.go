package dengta

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// AlibabaPicturesDengtaOrderStatusChange 天下秀订单状态变更通知
// alibaba.pictures.dengta.order.status.change
//
// 天下秀订单状态变更通知
func AlibabaPicturesDengtaOrderStatusChange(clt *core.SDKClient, req *dengta.AlibabaPicturesDengtaOrderStatusChangeAPIRequest, resp *dengta.AlibabaPicturesDengtaOrderStatusChangeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
