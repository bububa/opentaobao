package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmVoucherTemplateListAPIResponse 获取优惠券模版列表 API返回值
// alibaba.alsc.crm.voucher.template.list
//
// 获取优惠券模版列表
type AlibabaAlscCrmVoucherTemplateListAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmVoucherTemplateListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmVoucherTemplateListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmVoucherTemplateListAPIResponseModel).Reset()
}

// AlibabaAlscCrmVoucherTemplateListAPIResponseModel is 获取优惠券模版列表 成功返回结果
type AlibabaAlscCrmVoucherTemplateListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_voucher_template_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmVoucherTemplateListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmVoucherTemplateListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmVoucherTemplateListAPIResponse)
	},
}

// GetAlibabaAlscCrmVoucherTemplateListAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmVoucherTemplateListAPIResponse
func GetAlibabaAlscCrmVoucherTemplateListAPIResponse() *AlibabaAlscCrmVoucherTemplateListAPIResponse {
	return poolAlibabaAlscCrmVoucherTemplateListAPIResponse.Get().(*AlibabaAlscCrmVoucherTemplateListAPIResponse)
}

// ReleaseAlibabaAlscCrmVoucherTemplateListAPIResponse 将 AlibabaAlscCrmVoucherTemplateListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmVoucherTemplateListAPIResponse(v *AlibabaAlscCrmVoucherTemplateListAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmVoucherTemplateListAPIResponse.Put(v)
}
