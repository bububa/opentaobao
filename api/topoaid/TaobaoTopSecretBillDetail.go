package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoTopSecretBillDetail 服务商的商家解密账单详情查询
// taobao.top.secret.bill.detail
//
// 服务商的商家解密账单详情查询，仅对90天内的账单提供SLA保障。
func TaobaoTopSecretBillDetail(clt *core.SDKClient, req *topoaid.TaobaoTopSecretBillDetailAPIRequest, session string) (*topoaid.TaobaoTopSecretBillDetailAPIResponse, error) {
	var resp topoaid.TaobaoTopSecretBillDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
