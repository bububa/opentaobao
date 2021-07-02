package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceCoreInvUploadAPIResponse 发票中台-发票结果回传 API返回值
// alibaba.einvoice.core.inv.upload
//
// 发票回传接口适用于以下场景：
// ① 阿里发票平台向ISV提交原始发票申请，ISV开具发票成功后，基于申请ID(apply_id)回传发票至阿里发票平台进行归集与交付。
// ② 直接回传发票给阿里发票平台，进行归集，并交付给业务前台和用户。
type AlibabaEinvoiceCoreInvUploadAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceCoreInvUploadAPIResponseModel
}

// AlibabaEinvoiceCoreInvUploadAPIResponseModel is 发票中台-发票结果回传 成功返回结果
type AlibabaEinvoiceCoreInvUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_core_inv_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
