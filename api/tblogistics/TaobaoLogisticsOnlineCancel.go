package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaologisticsonlinecancel 取消物流订单接口
// taobao.logistics.online.cancel
//
// 调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。
func Taobaologisticsonlinecancel(clt *core.SDKClient, req *tblogistics.TaobaologisticsonlinecancelAPIRequest, session string) (*tblogistics.TaobaologisticsonlinecancelAPIResponse, error) {
	var resp tblogistics.TaobaologisticsonlinecancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
