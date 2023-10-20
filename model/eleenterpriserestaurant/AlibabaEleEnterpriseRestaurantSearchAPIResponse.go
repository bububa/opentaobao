package eleenterpriserestaurant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseRestaurantSearchAPIResponse 餐厅列表 API返回值
// alibaba.ele.enterprise.restaurant.search
//
// 餐厅列表
type AlibabaEleEnterpriseRestaurantSearchAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseRestaurantSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseRestaurantSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseRestaurantSearchAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseRestaurantSearchAPIResponseModel is 餐厅列表 成功返回结果
type AlibabaEleEnterpriseRestaurantSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_restaurant_search_response"`
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
func (m *AlibabaEleEnterpriseRestaurantSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
	m.EnterpriseData = nil
}

var poolAlibabaEleEnterpriseRestaurantSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseRestaurantSearchAPIResponse)
	},
}

// GetAlibabaEleEnterpriseRestaurantSearchAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseRestaurantSearchAPIResponse
func GetAlibabaEleEnterpriseRestaurantSearchAPIResponse() *AlibabaEleEnterpriseRestaurantSearchAPIResponse {
	return poolAlibabaEleEnterpriseRestaurantSearchAPIResponse.Get().(*AlibabaEleEnterpriseRestaurantSearchAPIResponse)
}

// ReleaseAlibabaEleEnterpriseRestaurantSearchAPIResponse 将 AlibabaEleEnterpriseRestaurantSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseRestaurantSearchAPIResponse(v *AlibabaEleEnterpriseRestaurantSearchAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseRestaurantSearchAPIResponse.Put(v)
}
