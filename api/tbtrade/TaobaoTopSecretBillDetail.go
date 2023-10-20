package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTopSecretBillDetail 服务商的商家解密账单详情查询
// taobao.top.secret.bill.detail
//
// 服务商的商家解密账单详情查询，仅对90天内的账单提供SLA保障。
func TaobaoTopSecretBillDetail(clt *core.SDKClient, req *tbtrade.TaobaoTopSecretBillDetailAPIRequest, resp *tbtrade.TaobaoTopSecretBillDetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
