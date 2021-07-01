package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

/* AlibabaAlscCrmOpenOrderBackflow
订单回流接口
alibaba.alsc.crm.open.order.backflow

回流isv订单接口 */
func AlibabaAlscCrmOpenOrderBackflow(clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenOrderBackflowAPIRequest, session string) (*alsc.AlibabaAlscCrmOpenOrderBackflowAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmOpenOrderBackflowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
