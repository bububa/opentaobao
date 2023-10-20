package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPricePromotionActivityQueryAPIResponse 查询盒马帮档期活动详情 API返回值
// alibaba.price.promotion.activity.query
//
// 查询盒马帮档期活动详情
type AlibabaPricePromotionActivityQueryAPIResponse struct {
	model.CommonResponse
	AlibabaPricePromotionActivityQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPricePromotionActivityQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPricePromotionActivityQueryAPIResponseModel).Reset()
}

// AlibabaPricePromotionActivityQueryAPIResponseModel is 查询盒马帮档期活动详情 成功返回结果
type AlibabaPricePromotionActivityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_price_promotion_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 大润发促销档期数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误参数
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 大润发档期数据
	TotalRecord int64 `json:"total_record,omitempty" xml:"total_record,omitempty"`
	// 接口调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPricePromotionActivityQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.Message = ""
	m.ResultCode = 0
	m.TotalRecord = 0
	m.IsSuccess = false
}

var poolAlibabaPricePromotionActivityQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPricePromotionActivityQueryAPIResponse)
	},
}

// GetAlibabaPricePromotionActivityQueryAPIResponse 从 sync.Pool 获取 AlibabaPricePromotionActivityQueryAPIResponse
func GetAlibabaPricePromotionActivityQueryAPIResponse() *AlibabaPricePromotionActivityQueryAPIResponse {
	return poolAlibabaPricePromotionActivityQueryAPIResponse.Get().(*AlibabaPricePromotionActivityQueryAPIResponse)
}

// ReleaseAlibabaPricePromotionActivityQueryAPIResponse 将 AlibabaPricePromotionActivityQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPricePromotionActivityQueryAPIResponse(v *AlibabaPricePromotionActivityQueryAPIResponse) {
	v.Reset()
	poolAlibabaPricePromotionActivityQueryAPIResponse.Put(v)
}
