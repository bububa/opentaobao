package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterSpserviceorderQueryAPIResponse 服务供应链服务单查询 API返回值
// alibaba.servicecenter.spserviceorder.query
//
// 服务供应链服务单查询，服务商通过此接口拉取用户的购买的服务信息，以此为依据为用户提供安装维修等服务
type AlibabaServicecenterSpserviceorderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterSpserviceorderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaServicecenterSpserviceorderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaServicecenterSpserviceorderQueryAPIResponseModel).Reset()
}

// AlibabaServicecenterSpserviceorderQueryAPIResponseModel is 服务供应链服务单查询 成功返回结果
type AlibabaServicecenterSpserviceorderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_spserviceorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *AlibabaServicecenterSpserviceorderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaServicecenterSpserviceorderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaServicecenterSpserviceorderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterSpserviceorderQueryAPIResponse)
	},
}

// GetAlibabaServicecenterSpserviceorderQueryAPIResponse 从 sync.Pool 获取 AlibabaServicecenterSpserviceorderQueryAPIResponse
func GetAlibabaServicecenterSpserviceorderQueryAPIResponse() *AlibabaServicecenterSpserviceorderQueryAPIResponse {
	return poolAlibabaServicecenterSpserviceorderQueryAPIResponse.Get().(*AlibabaServicecenterSpserviceorderQueryAPIResponse)
}

// ReleaseAlibabaServicecenterSpserviceorderQueryAPIResponse 将 AlibabaServicecenterSpserviceorderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaServicecenterSpserviceorderQueryAPIResponse(v *AlibabaServicecenterSpserviceorderQueryAPIResponse) {
	v.Reset()
	poolAlibabaServicecenterSpserviceorderQueryAPIResponse.Put(v)
}
