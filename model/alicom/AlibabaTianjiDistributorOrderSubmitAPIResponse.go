package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianjiDistributorOrderSubmitAPIResponse 分销商提交受理订单 API返回值
// alibaba.tianji.distributor.order.submit
//
// 分销商提交受理订单，如合约订购、充值受理等
type AlibabaTianjiDistributorOrderSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaTianjiDistributorOrderSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTianjiDistributorOrderSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTianjiDistributorOrderSubmitAPIResponseModel).Reset()
}

// AlibabaTianjiDistributorOrderSubmitAPIResponseModel is 分销商提交受理订单 成功返回结果
type AlibabaTianjiDistributorOrderSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianji_distributor_order_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTianjiDistributorOrderSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTianjiDistributorOrderSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTianjiDistributorOrderSubmitAPIResponse)
	},
}

// GetAlibabaTianjiDistributorOrderSubmitAPIResponse 从 sync.Pool 获取 AlibabaTianjiDistributorOrderSubmitAPIResponse
func GetAlibabaTianjiDistributorOrderSubmitAPIResponse() *AlibabaTianjiDistributorOrderSubmitAPIResponse {
	return poolAlibabaTianjiDistributorOrderSubmitAPIResponse.Get().(*AlibabaTianjiDistributorOrderSubmitAPIResponse)
}

// ReleaseAlibabaTianjiDistributorOrderSubmitAPIResponse 将 AlibabaTianjiDistributorOrderSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTianjiDistributorOrderSubmitAPIResponse(v *AlibabaTianjiDistributorOrderSubmitAPIResponse) {
	v.Reset()
	poolAlibabaTianjiDistributorOrderSubmitAPIResponse.Put(v)
}
