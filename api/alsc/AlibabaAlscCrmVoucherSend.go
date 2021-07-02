package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmVoucherSend 发送券给指定用户
// alibaba.alsc.crm.voucher.send
//
// 发送券给指定用户
func AlibabaAlscCrmVoucherSend(clt *core.SDKClient, req *alsc.AlibabaAlscCrmVoucherSendAPIRequest, session string) (*alsc.AlibabaAlscCrmVoucherSendAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmVoucherSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
