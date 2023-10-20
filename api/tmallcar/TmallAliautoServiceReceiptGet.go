package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoServiceReceiptGet isv查询服务工单详情
// tmall.aliauto.service.receipt.get
//
// isv查询服务工单详情
func TmallAliautoServiceReceiptGet(clt *core.SDKClient, req *tmallcar.TmallAliautoServiceReceiptGetAPIRequest, resp *tmallcar.TmallAliautoServiceReceiptGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
