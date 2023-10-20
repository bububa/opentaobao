package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmvouchersend 发送券给指定用户
// alibaba.alsc.crm.voucher.send
//
// 发送券给指定用户
func Alibabaalsccrmvouchersend(clt *core.SDKClient, req *alsc.AlibabaalsccrmvouchersendAPIRequest, session string) (*alsc.AlibabaalsccrmvouchersendAPIResponse, error) {
	var resp alsc.AlibabaalsccrmvouchersendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
