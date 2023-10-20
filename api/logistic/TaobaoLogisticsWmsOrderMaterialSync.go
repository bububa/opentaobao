package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticswmsordermaterialsync 仓服务商订单包材耗材信息同步
// taobao.logistics.wms.order.material.sync
//
// 仓服务商订单包材耗材信息同步
func Taobaologisticswmsordermaterialsync(clt *core.SDKClient, req *logistic.TaobaologisticswmsordermaterialsyncAPIRequest, session string) (*logistic.TaobaologisticswmsordermaterialsyncAPIResponse, error) {
	var resp logistic.TaobaologisticswmsordermaterialsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
