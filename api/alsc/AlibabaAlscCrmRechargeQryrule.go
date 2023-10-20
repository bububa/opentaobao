package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrechargeqryrule 储值规则下行
// alibaba.alsc.crm.recharge.qryrule
//
// 储值规则下行
func Alibabaalsccrmrechargeqryrule(clt *core.SDKClient, req *alsc.AlibabaalsccrmrechargeqryruleAPIRequest, session string) (*alsc.AlibabaalsccrmrechargeqryruleAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrechargeqryruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
