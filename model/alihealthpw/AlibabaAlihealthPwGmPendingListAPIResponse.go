package alihealthpw

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwGmPendingListAPIResponse 同情用药待审核工单查询接口 API返回值
// alibaba.alihealth.pw.gm.pending.list
//
// 同情用药待审核工单查询接口，提供给合作方用来查询待处理工单列表
type AlibabaAlihealthPwGmPendingListAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPwGmPendingListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwGmPendingListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthPwGmPendingListAPIResponseModel).Reset()
}

// AlibabaAlihealthPwGmPendingListAPIResponseModel is 同情用药待审核工单查询接口 成功返回结果
type AlibabaAlihealthPwGmPendingListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_gm_pending_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwGmPendingListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthPwGmPendingListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPwGmPendingListAPIResponse)
	},
}

// GetAlibabaAlihealthPwGmPendingListAPIResponse 从 sync.Pool 获取 AlibabaAlihealthPwGmPendingListAPIResponse
func GetAlibabaAlihealthPwGmPendingListAPIResponse() *AlibabaAlihealthPwGmPendingListAPIResponse {
	return poolAlibabaAlihealthPwGmPendingListAPIResponse.Get().(*AlibabaAlihealthPwGmPendingListAPIResponse)
}

// ReleaseAlibabaAlihealthPwGmPendingListAPIResponse 将 AlibabaAlihealthPwGmPendingListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthPwGmPendingListAPIResponse(v *AlibabaAlihealthPwGmPendingListAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthPwGmPendingListAPIResponse.Put(v)
}
