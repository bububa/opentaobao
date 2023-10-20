package alihealthpw

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwGmDetailAPIResponse 同情用药申请单详情接口 API返回值
// alibaba.alihealth.pw.gm.detail
//
// 同情用药申请单详情接口，提供给合作方查询申请单详情
type AlibabaAlihealthPwGmDetailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPwGmDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwGmDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthPwGmDetailAPIResponseModel).Reset()
}

// AlibabaAlihealthPwGmDetailAPIResponseModel is 同情用药申请单详情接口 成功返回结果
type AlibabaAlihealthPwGmDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_gm_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwGmDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthPwGmDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPwGmDetailAPIResponse)
	},
}

// GetAlibabaAlihealthPwGmDetailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthPwGmDetailAPIResponse
func GetAlibabaAlihealthPwGmDetailAPIResponse() *AlibabaAlihealthPwGmDetailAPIResponse {
	return poolAlibabaAlihealthPwGmDetailAPIResponse.Get().(*AlibabaAlihealthPwGmDetailAPIResponse)
}

// ReleaseAlibabaAlihealthPwGmDetailAPIResponse 将 AlibabaAlihealthPwGmDetailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthPwGmDetailAPIResponse(v *AlibabaAlihealthPwGmDetailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthPwGmDetailAPIResponse.Put(v)
}
