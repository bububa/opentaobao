package cainiaolocker

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoEndpointLockerTopWithholdQueryAPIResponse 查询能否代扣 API返回值
// cainiao.endpoint.locker.top.withhold.query
//
// 查询是否有代扣欠款，是否签署代扣协议。
type CainiaoEndpointLockerTopWithholdQueryAPIResponse struct {
	model.CommonResponse
	CainiaoEndpointLockerTopWithholdQueryAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoEndpointLockerTopWithholdQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoEndpointLockerTopWithholdQueryAPIResponseModel).Reset()
}

// CainiaoEndpointLockerTopWithholdQueryAPIResponseModel is 查询能否代扣 成功返回结果
type CainiaoEndpointLockerTopWithholdQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_endpoint_locker_top_withhold_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// response
	Result *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoEndpointLockerTopWithholdQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoEndpointLockerTopWithholdQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoEndpointLockerTopWithholdQueryAPIResponse)
	},
}

// GetCainiaoEndpointLockerTopWithholdQueryAPIResponse 从 sync.Pool 获取 CainiaoEndpointLockerTopWithholdQueryAPIResponse
func GetCainiaoEndpointLockerTopWithholdQueryAPIResponse() *CainiaoEndpointLockerTopWithholdQueryAPIResponse {
	return poolCainiaoEndpointLockerTopWithholdQueryAPIResponse.Get().(*CainiaoEndpointLockerTopWithholdQueryAPIResponse)
}

// ReleaseCainiaoEndpointLockerTopWithholdQueryAPIResponse 将 CainiaoEndpointLockerTopWithholdQueryAPIResponse 保存到 sync.Pool
func ReleaseCainiaoEndpointLockerTopWithholdQueryAPIResponse(v *CainiaoEndpointLockerTopWithholdQueryAPIResponse) {
	v.Reset()
	poolCainiaoEndpointLockerTopWithholdQueryAPIResponse.Put(v)
}
