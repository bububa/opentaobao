package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceapplygetAPIResponse 开票申请数据获取接口 API返回值
// alibaba.einvoice.apply.get
//
// ERP获取开票申请数据
type AlibabaeinvoiceapplygetAPIResponse struct {
	model.CommonResponse
	AlibabaeinvoiceapplygetAPIResponseModel
}

// AlibabaeinvoiceapplygetAPIResponseModel is 开票申请数据获取接口 成功返回结果
type AlibabaeinvoiceapplygetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_apply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 开票明细
	ApplyList []Apply `json:"apply_list,omitempty" xml:"apply_list>apply,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
