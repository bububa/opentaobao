package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcustomervoucherlistAPIResponse 获取顾客优惠券列表 API返回值
// alibaba.alsc.crm.customer.voucher.list
//
// 获取顾客优惠券列表
type AlibabaalsccrmcustomervoucherlistAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmcustomervoucherlistAPIResponseModel
}

// AlibabaalsccrmcustomervoucherlistAPIResponseModel is 获取顾客优惠券列表 成功返回结果
type AlibabaalsccrmcustomervoucherlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_customer_voucher_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
