package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaologisticswmspackageentryorderconfirm 包裹入库单确认
// taobao.logistics.wms.packageentryorder.confirm
//
// 包裹入库单确认
func Taobaologisticswmspackageentryorderconfirm(clt *core.SDKClient, req *tblogistics.TaobaologisticswmspackageentryorderconfirmAPIRequest, session string) (*tblogistics.TaobaologisticswmspackageentryorderconfirmAPIResponse, error) {
	var resp tblogistics.TaobaologisticswmspackageentryorderconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
