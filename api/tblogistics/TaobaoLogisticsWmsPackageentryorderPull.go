package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaologisticswmspackageentryorderpull 包裹入库单拉单
// taobao.logistics.wms.packageentryorder.pull
//
// 包裹入库单拉单
func Taobaologisticswmspackageentryorderpull(clt *core.SDKClient, req *tblogistics.TaobaologisticswmspackageentryorderpullAPIRequest, session string) (*tblogistics.TaobaologisticswmspackageentryorderpullAPIResponse, error) {
	var resp tblogistics.TaobaologisticswmspackageentryorderpullAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
