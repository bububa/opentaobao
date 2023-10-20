package eleenterpriserestaurant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseRestaurantMixAPIResponse 混合搜索店铺 API返回值
// alibaba.ele.enterprise.restaurant.mix
//
// 混合搜索店铺
type AlibabaEleEnterpriseRestaurantMixAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseRestaurantMixAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseRestaurantMixAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseRestaurantMixAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseRestaurantMixAPIResponseModel is 混合搜索店铺 成功返回结果
type AlibabaEleEnterpriseRestaurantMixAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_restaurant_mix_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
	// 返回值信息
	EnterpriseData *EnterpriseData `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseRestaurantMixAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
	m.EnterpriseData = nil
}

var poolAlibabaEleEnterpriseRestaurantMixAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseRestaurantMixAPIResponse)
	},
}

// GetAlibabaEleEnterpriseRestaurantMixAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseRestaurantMixAPIResponse
func GetAlibabaEleEnterpriseRestaurantMixAPIResponse() *AlibabaEleEnterpriseRestaurantMixAPIResponse {
	return poolAlibabaEleEnterpriseRestaurantMixAPIResponse.Get().(*AlibabaEleEnterpriseRestaurantMixAPIResponse)
}

// ReleaseAlibabaEleEnterpriseRestaurantMixAPIResponse 将 AlibabaEleEnterpriseRestaurantMixAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseRestaurantMixAPIResponse(v *AlibabaEleEnterpriseRestaurantMixAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseRestaurantMixAPIResponse.Put(v)
}
