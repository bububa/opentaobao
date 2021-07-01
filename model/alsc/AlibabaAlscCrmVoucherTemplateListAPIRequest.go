package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmVoucherTemplateListAPIRequest
获取优惠券模版列表 API请求
alibaba.alsc.crm.voucher.template.list

获取优惠券模版列表 */
type AlibabaAlscCrmVoucherTemplateListAPIRequest struct {
	model.Params
	// 获取优惠模版规则请求参数
	_voucherTemplateOpenReq *VoucherTemplateOpenReq
}

// New
