package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointGetrefAPIResponse 查询业务实体关联口径2 API返回值
// alibaba.legal.standpoint.getref
//
// 口径查询
type AlibabaLegalStandpointGetrefAPIResponse struct {
	model.CommonResponse
	AlibabaLegalStandpointGetrefAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointGetrefAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalStandpointGetrefAPIResponseModel).Reset()
}

// AlibabaLegalStandpointGetrefAPIResponseModel is 查询业务实体关联口径2 成功返回结果
type AlibabaLegalStandpointGetrefAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_getref_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统错误
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回状态
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 分页对象
	Content *Page `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointGetrefAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ErrorCodeRes = 0
	m.Content = nil
	m.SuccessRes = false
}

var poolAlibabaLegalStandpointGetrefAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalStandpointGetrefAPIResponse)
	},
}

// GetAlibabaLegalStandpointGetrefAPIResponse 从 sync.Pool 获取 AlibabaLegalStandpointGetrefAPIResponse
func GetAlibabaLegalStandpointGetrefAPIResponse() *AlibabaLegalStandpointGetrefAPIResponse {
	return poolAlibabaLegalStandpointGetrefAPIResponse.Get().(*AlibabaLegalStandpointGetrefAPIResponse)
}

// ReleaseAlibabaLegalStandpointGetrefAPIResponse 将 AlibabaLegalStandpointGetrefAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalStandpointGetrefAPIResponse(v *AlibabaLegalStandpointGetrefAPIResponse) {
	v.Reset()
	poolAlibabaLegalStandpointGetrefAPIResponse.Put(v)
}
