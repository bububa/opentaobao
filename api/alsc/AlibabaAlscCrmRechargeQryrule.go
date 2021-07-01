package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

/* AlibabaAlscCrmRechargeQryrule
储值规则下行
alibaba.alsc.crm.recharge.qryrule

储值规则下行 */
func AlibabaAlscCrmRechargeQryrule(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeQryruleAPIRequest, session string) (*alsc.AlibabaAlscCrmRechargeQryruleAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmRechargeQryruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
