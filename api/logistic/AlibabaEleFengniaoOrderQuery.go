package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoOrderQuery 查询订单基本信息
// alibaba.ele.fengniao.order.query
//
// 查询订单基本信息
func AlibabaEleFengniaoOrderQuery(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoOrderQueryAPIRequest, resp *logistic.AlibabaEleFengniaoOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
