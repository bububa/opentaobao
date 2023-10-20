package dengta

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// AlibabaPicturesDengtaImsOrderStatusChange 天下秀回传订单执行状态变动
// alibaba.pictures.dengta.ims.order.status.change
//
// 天下秀回传订单执行状态变动
func AlibabaPicturesDengtaImsOrderStatusChange(clt *core.SDKClient, req *dengta.AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest, resp *dengta.AlibabaPicturesDengtaImsOrderStatusChangeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
