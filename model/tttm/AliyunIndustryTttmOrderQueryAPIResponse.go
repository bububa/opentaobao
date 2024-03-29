package tttm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunIndustryTttmOrderQueryAPIResponse 天天特卖数字工厂订单获取 API返回值
// aliyun.industry.tttm.order.query
//
// 获取阿里云数字工厂内天天特卖业务的订单
type AliyunIndustryTttmOrderQueryAPIResponse struct {
	model.CommonResponse
	AliyunIndustryTttmOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunIndustryTttmOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunIndustryTttmOrderQueryAPIResponseModel).Reset()
}

// AliyunIndustryTttmOrderQueryAPIResponseModel is 天天特卖数字工厂订单获取 成功返回结果
type AliyunIndustryTttmOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_industry_tttm_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单
	Result *OrderDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliyunIndustryTttmOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliyunIndustryTttmOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunIndustryTttmOrderQueryAPIResponse)
	},
}

// GetAliyunIndustryTttmOrderQueryAPIResponse 从 sync.Pool 获取 AliyunIndustryTttmOrderQueryAPIResponse
func GetAliyunIndustryTttmOrderQueryAPIResponse() *AliyunIndustryTttmOrderQueryAPIResponse {
	return poolAliyunIndustryTttmOrderQueryAPIResponse.Get().(*AliyunIndustryTttmOrderQueryAPIResponse)
}

// ReleaseAliyunIndustryTttmOrderQueryAPIResponse 将 AliyunIndustryTttmOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAliyunIndustryTttmOrderQueryAPIResponse(v *AliyunIndustryTttmOrderQueryAPIResponse) {
	v.Reset()
	poolAliyunIndustryTttmOrderQueryAPIResponse.Put(v)
}
