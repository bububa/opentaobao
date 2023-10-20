package alihealthpw

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwGmIdsListAPIResponse 同情用药根据申请单列表查询申请单 API返回值
// alibaba.alihealth.pw.gm.ids.list
//
// 同情用药根据申请单列表查询申请单
type AlibabaAlihealthPwGmIdsListAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPwGmIdsListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwGmIdsListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthPwGmIdsListAPIResponseModel).Reset()
}

// AlibabaAlihealthPwGmIdsListAPIResponseModel is 同情用药根据申请单列表查询申请单 成功返回结果
type AlibabaAlihealthPwGmIdsListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_gm_ids_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwGmIdsListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthPwGmIdsListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPwGmIdsListAPIResponse)
	},
}

// GetAlibabaAlihealthPwGmIdsListAPIResponse 从 sync.Pool 获取 AlibabaAlihealthPwGmIdsListAPIResponse
func GetAlibabaAlihealthPwGmIdsListAPIResponse() *AlibabaAlihealthPwGmIdsListAPIResponse {
	return poolAlibabaAlihealthPwGmIdsListAPIResponse.Get().(*AlibabaAlihealthPwGmIdsListAPIResponse)
}

// ReleaseAlibabaAlihealthPwGmIdsListAPIResponse 将 AlibabaAlihealthPwGmIdsListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthPwGmIdsListAPIResponse(v *AlibabaAlihealthPwGmIdsListAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthPwGmIdsListAPIResponse.Put(v)
}
