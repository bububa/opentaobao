package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPricePromotionCreateAPIResponse 营销档期活动创建 API返回值
// alibaba.price.promotion.create
//
// 大润发-盒马帮提供新增创建营销活动
type AlibabaPricePromotionCreateAPIResponse struct {
	model.CommonResponse
	AlibabaPricePromotionCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPricePromotionCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPricePromotionCreateAPIResponseModel).Reset()
}

// AlibabaPricePromotionCreateAPIResponseModel is 营销档期活动创建 成功返回结果
type AlibabaPricePromotionCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_price_promotion_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 错误编码，本期不作识别
	SystemCode string `json:"system_code,omitempty" xml:"system_code,omitempty"`
	// 档期活动ID
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 数量，本期不启用
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 创建是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPricePromotionCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorDesc = ""
	m.SystemCode = ""
	m.Result = 0
	m.TotalNum = 0
	m.IsSuccess = false
}

var poolAlibabaPricePromotionCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPricePromotionCreateAPIResponse)
	},
}

// GetAlibabaPricePromotionCreateAPIResponse 从 sync.Pool 获取 AlibabaPricePromotionCreateAPIResponse
func GetAlibabaPricePromotionCreateAPIResponse() *AlibabaPricePromotionCreateAPIResponse {
	return poolAlibabaPricePromotionCreateAPIResponse.Get().(*AlibabaPricePromotionCreateAPIResponse)
}

// ReleaseAlibabaPricePromotionCreateAPIResponse 将 AlibabaPricePromotionCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPricePromotionCreateAPIResponse(v *AlibabaPricePromotionCreateAPIResponse) {
	v.Reset()
	poolAlibabaPricePromotionCreateAPIResponse.Put(v)
}
