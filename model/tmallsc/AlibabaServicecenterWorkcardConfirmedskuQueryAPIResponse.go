package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse 查询确认履行的服务项 API返回值
// alibaba.servicecenter.workcard.confirmedsku.query
//
// 查询确认履行的服务项
type AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponseModel).Reset()
}

// AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponseModel is 查询确认履行的服务项 成功返回结果
type AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_workcard_confirmedsku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaServicecenterWorkcardConfirmedskuQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse)
	},
}

// GetAlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse 从 sync.Pool 获取 AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse
func GetAlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse() *AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse {
	return poolAlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse.Get().(*AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse)
}

// ReleaseAlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse 将 AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse(v *AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse) {
	v.Reset()
	poolAlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse.Put(v)
}
