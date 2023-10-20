package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaologisticswmspackagedeliveryorderconfirm 包裹出库单确认
// taobao.logistics.wms.packagedeliveryorder.confirm
//
// 包裹出库单确认
func Taobaologisticswmspackagedeliveryorderconfirm(clt *core.SDKClient, req *tblogistics.TaobaologisticswmspackagedeliveryorderconfirmAPIRequest, session string) (*tblogistics.TaobaologisticswmspackagedeliveryorderconfirmAPIResponse, error) {
	var resp tblogistics.TaobaologisticswmspackagedeliveryorderconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
