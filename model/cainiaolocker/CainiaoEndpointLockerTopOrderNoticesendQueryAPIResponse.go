package cainiaolocker

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse 查询订单是否由裹裹发送消息 API返回值
// cainiao.endpoint.locker.top.order.noticesend.query
//
// 合作公司查询消息发送的接口，判断是否裹裹发送消息
type CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse struct {
	model.CommonResponse
	CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponseModel).Reset()
}

// CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponseModel is 查询订单是否由裹裹发送消息 成功返回结果
type CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_endpoint_locker_top_order_noticesend_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse)
	},
}

// GetCainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse 从 sync.Pool 获取 CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse
func GetCainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse() *CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse {
	return poolCainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse.Get().(*CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse)
}

// ReleaseCainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse 将 CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse 保存到 sync.Pool
func ReleaseCainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse(v *CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse) {
	v.Reset()
	poolCainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse.Put(v)
}
