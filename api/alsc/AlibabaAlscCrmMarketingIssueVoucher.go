package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmmarketingissuevoucher 发券
// alibaba.alsc.crm.marketing.issue.voucher
//
// 提供发券功能
func Alibabaalsccrmmarketingissuevoucher(clt *core.SDKClient, req *alsc.AlibabaalsccrmmarketingissuevoucherAPIRequest, session string) (*alsc.AlibabaalsccrmmarketingissuevoucherAPIResponse, error) {
	var resp alsc.AlibabaalsccrmmarketingissuevoucherAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
