package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoReceiptStateUpdate 服务工单状态更新
// tmall.aliauto.receipt.state.update
//
// 二轮车服务工单状态更新
func TmallAliautoReceiptStateUpdate(clt *core.SDKClient, req *tmallcar.TmallAliautoReceiptStateUpdateAPIRequest, resp *tmallcar.TmallAliautoReceiptStateUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
