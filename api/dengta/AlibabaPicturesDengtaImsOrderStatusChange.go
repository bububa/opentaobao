package dengta

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// AlibabaPicturesDengtaImsOrderStatusChange 天下秀回传订单执行状态变动
// alibaba.pictures.dengta.ims.order.status.change
//
// 天下秀回传订单执行状态变动
func AlibabaPicturesDengtaImsOrderStatusChange(clt *core.SDKClient, req *dengta.AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest, session string) (*dengta.AlibabaPicturesDengtaImsOrderStatusChangeAPIResponse, error) {
	var resp dengta.AlibabaPicturesDengtaImsOrderStatusChangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
