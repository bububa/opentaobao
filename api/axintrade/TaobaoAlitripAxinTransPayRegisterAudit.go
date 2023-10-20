package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripAxinTransPayRegisterAudit 阿信支付入驻审核通知
// taobao.alitrip.axin.trans.pay.register.audit
//
// 阿信支付入驻审核通知
func TaobaoAlitripAxinTransPayRegisterAudit(clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransPayRegisterAuditAPIRequest, session string) (*axintrade.TaobaoAlitripAxinTransPayRegisterAuditAPIResponse, error) {
	var resp axintrade.TaobaoAlitripAxinTransPayRegisterAuditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
