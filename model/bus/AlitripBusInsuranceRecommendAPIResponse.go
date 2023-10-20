package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBusInsuranceRecommendAPIResponse 汽车票保险推荐接口 API返回值
// alitrip.bus.insurance.recommend
//
// 汽车票保险推荐接口-供商家售卖飞猪保险使用
type AlitripBusInsuranceRecommendAPIResponse struct {
	model.CommonResponse
	AlitripBusInsuranceRecommendAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBusInsuranceRecommendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBusInsuranceRecommendAPIResponseModel).Reset()
}

// AlitripBusInsuranceRecommendAPIResponseModel is 汽车票保险推荐接口 成功返回结果
type AlitripBusInsuranceRecommendAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_bus_insurance_recommend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 保险产品列表
	InsProductList []InsuranceProductVo `json:"ins_product_list,omitempty" xml:"ins_product_list>insurance_product_vo,omitempty"`
	// 接口业务成功时为空
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// true:成功；false:失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBusInsuranceRecommendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InsProductList = m.InsProductList[:0]
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolAlitripBusInsuranceRecommendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBusInsuranceRecommendAPIResponse)
	},
}

// GetAlitripBusInsuranceRecommendAPIResponse 从 sync.Pool 获取 AlitripBusInsuranceRecommendAPIResponse
func GetAlitripBusInsuranceRecommendAPIResponse() *AlitripBusInsuranceRecommendAPIResponse {
	return poolAlitripBusInsuranceRecommendAPIResponse.Get().(*AlitripBusInsuranceRecommendAPIResponse)
}

// ReleaseAlitripBusInsuranceRecommendAPIResponse 将 AlitripBusInsuranceRecommendAPIResponse 保存到 sync.Pool
func ReleaseAlitripBusInsuranceRecommendAPIResponse(v *AlitripBusInsuranceRecommendAPIResponse) {
	v.Reset()
	poolAlitripBusInsuranceRecommendAPIResponse.Put(v)
}
