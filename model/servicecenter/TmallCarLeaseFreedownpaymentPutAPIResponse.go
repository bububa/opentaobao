package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseFreedownpaymentPutAPIResponse 同步直租车免首付商品活动信息 API返回值
// tmall.car.lease.freedownpayment.put
//
// 汽车行业直租车免首付需求中，用与对商品打标，活动范围设置，在消费者端商品详情页、订单等环节透出，表示该商品为直租免首付商品。
type TmallCarLeaseFreedownpaymentPutAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseFreedownpaymentPutAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarLeaseFreedownpaymentPutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarLeaseFreedownpaymentPutAPIResponseModel).Reset()
}

// TmallCarLeaseFreedownpaymentPutAPIResponseModel is 同步直租车免首付商品活动信息 成功返回结果
type TmallCarLeaseFreedownpaymentPutAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_freedownpayment_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallCarLeaseFreedownpaymentPutResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarLeaseFreedownpaymentPutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarLeaseFreedownpaymentPutAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarLeaseFreedownpaymentPutAPIResponse)
	},
}

// GetTmallCarLeaseFreedownpaymentPutAPIResponse 从 sync.Pool 获取 TmallCarLeaseFreedownpaymentPutAPIResponse
func GetTmallCarLeaseFreedownpaymentPutAPIResponse() *TmallCarLeaseFreedownpaymentPutAPIResponse {
	return poolTmallCarLeaseFreedownpaymentPutAPIResponse.Get().(*TmallCarLeaseFreedownpaymentPutAPIResponse)
}

// ReleaseTmallCarLeaseFreedownpaymentPutAPIResponse 将 TmallCarLeaseFreedownpaymentPutAPIResponse 保存到 sync.Pool
func ReleaseTmallCarLeaseFreedownpaymentPutAPIResponse(v *TmallCarLeaseFreedownpaymentPutAPIResponse) {
	v.Reset()
	poolTmallCarLeaseFreedownpaymentPutAPIResponse.Put(v)
}
