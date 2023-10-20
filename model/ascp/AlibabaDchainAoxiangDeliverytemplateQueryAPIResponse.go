package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangDeliverytemplateQueryAPIResponse 商家运费模板查询 API返回值
// alibaba.dchain.aoxiang.deliverytemplate.query
//
// 商家运费模板查询
type AlibabaDchainAoxiangDeliverytemplateQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangDeliverytemplateQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangDeliverytemplateQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangDeliverytemplateQueryAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangDeliverytemplateQueryAPIResponseModel is 商家运费模板查询 成功返回结果
type AlibabaDchainAoxiangDeliverytemplateQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_deliverytemplate_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	QueryDeliverytemplateResponse *TopResponse `json:"query_deliverytemplate_response,omitempty" xml:"query_deliverytemplate_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangDeliverytemplateQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.QueryDeliverytemplateResponse = nil
}

var poolAlibabaDchainAoxiangDeliverytemplateQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangDeliverytemplateQueryAPIResponse)
	},
}

// GetAlibabaDchainAoxiangDeliverytemplateQueryAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangDeliverytemplateQueryAPIResponse
func GetAlibabaDchainAoxiangDeliverytemplateQueryAPIResponse() *AlibabaDchainAoxiangDeliverytemplateQueryAPIResponse {
	return poolAlibabaDchainAoxiangDeliverytemplateQueryAPIResponse.Get().(*AlibabaDchainAoxiangDeliverytemplateQueryAPIResponse)
}

// ReleaseAlibabaDchainAoxiangDeliverytemplateQueryAPIResponse 将 AlibabaDchainAoxiangDeliverytemplateQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangDeliverytemplateQueryAPIResponse(v *AlibabaDchainAoxiangDeliverytemplateQueryAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangDeliverytemplateQueryAPIResponse.Put(v)
}
