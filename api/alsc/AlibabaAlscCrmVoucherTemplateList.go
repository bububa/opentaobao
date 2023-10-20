package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmVoucherTemplateList 获取优惠券模版列表
// alibaba.alsc.crm.voucher.template.list
//
// 获取优惠券模版列表
func AlibabaAlscCrmVoucherTemplateList(clt *core.SDKClient, req *alsc.AlibabaAlscCrmVoucherTemplateListAPIRequest, session string) (*alsc.AlibabaAlscCrmVoucherTemplateListAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmVoucherTemplateListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
