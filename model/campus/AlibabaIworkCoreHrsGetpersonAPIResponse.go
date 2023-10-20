package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIworkCoreHrsGetpersonAPIResponse 获取神鲸用户基本信息 API返回值
// alibaba.iwork.core.hrs.getperson
//
// 神鲸用户的基本信息查询，根据PERSON_ID或者用户ACCOUNT_ID查询
type AlibabaIworkCoreHrsGetpersonAPIResponse struct {
	model.CommonResponse
	AlibabaIworkCoreHrsGetpersonAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIworkCoreHrsGetpersonAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIworkCoreHrsGetpersonAPIResponseModel).Reset()
}

// AlibabaIworkCoreHrsGetpersonAPIResponseModel is 获取神鲸用户基本信息 成功返回结果
type AlibabaIworkCoreHrsGetpersonAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_iwork_core_hrs_getperson_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIworkCoreHrsGetpersonAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIworkCoreHrsGetpersonAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIworkCoreHrsGetpersonAPIResponse)
	},
}

// GetAlibabaIworkCoreHrsGetpersonAPIResponse 从 sync.Pool 获取 AlibabaIworkCoreHrsGetpersonAPIResponse
func GetAlibabaIworkCoreHrsGetpersonAPIResponse() *AlibabaIworkCoreHrsGetpersonAPIResponse {
	return poolAlibabaIworkCoreHrsGetpersonAPIResponse.Get().(*AlibabaIworkCoreHrsGetpersonAPIResponse)
}

// ReleaseAlibabaIworkCoreHrsGetpersonAPIResponse 将 AlibabaIworkCoreHrsGetpersonAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIworkCoreHrsGetpersonAPIResponse(v *AlibabaIworkCoreHrsGetpersonAPIResponse) {
	v.Reset()
	poolAlibabaIworkCoreHrsGetpersonAPIResponse.Put(v)
}
