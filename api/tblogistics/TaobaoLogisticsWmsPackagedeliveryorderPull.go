package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaologisticswmspackagedeliveryorderpull 包裹出库单拉单
// taobao.logistics.wms.packagedeliveryorder.pull
//
// 包裹出库单拉单
func Taobaologisticswmspackagedeliveryorderpull(clt *core.SDKClient, req *tblogistics.TaobaologisticswmspackagedeliveryorderpullAPIRequest, session string) (*tblogistics.TaobaologisticswmspackagedeliveryorderpullAPIResponse, error) {
	var resp tblogistics.TaobaologisticswmspackagedeliveryorderpullAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
