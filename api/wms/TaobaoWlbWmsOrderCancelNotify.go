package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsOrderCancelNotify 单据取消接口
// taobao.wlb.wms.order.cancel.notify
//
// 单据取消接口
func TaobaoWlbWmsOrderCancelNotify(clt *core.SDKClient, req *wms.TaobaoWlbWmsOrderCancelNotifyAPIRequest, session string) (*wms.TaobaoWlbWmsOrderCancelNotifyAPIResponse, error) {
	var resp wms.TaobaoWlbWmsOrderCancelNotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
