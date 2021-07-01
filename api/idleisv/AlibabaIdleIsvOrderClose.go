package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

/* AlibabaIdleIsvOrderClose
服务商闲鱼卖家主动关闭订单
alibaba.idle.isv.order.close

供外部服务商 isv 提供卖家主动关闭交易订单的功能 */
func AlibabaIdleIsvOrderClose(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvOrderCloseAPIRequest, session string) (*idleisv.AlibabaIdleIsvOrderCloseAPIResponse, error) {
	var resp idleisv.AlibabaIdleIsvOrderCloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
