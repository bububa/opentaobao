package eleenterpriserestaurant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseRestaurantGetAPIResponse 查询餐厅信息 API返回值
// alibaba.ele.enterprise.restaurant.get
//
// 查询餐厅信息
type AlibabaEleEnterpriseRestaurantGetAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseRestaurantGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseRestaurantGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseRestaurantGetAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseRestaurantGetAPIResponseModel is 查询餐厅信息 成功返回结果
type AlibabaEleEnterpriseRestaurantGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_restaurant_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
	// 返回餐厅信息
	EnterpriseData *EnterpriseData `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseRestaurantGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
	m.EnterpriseData = nil
}

var poolAlibabaEleEnterpriseRestaurantGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseRestaurantGetAPIResponse)
	},
}

// GetAlibabaEleEnterpriseRestaurantGetAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseRestaurantGetAPIResponse
func GetAlibabaEleEnterpriseRestaurantGetAPIResponse() *AlibabaEleEnterpriseRestaurantGetAPIResponse {
	return poolAlibabaEleEnterpriseRestaurantGetAPIResponse.Get().(*AlibabaEleEnterpriseRestaurantGetAPIResponse)
}

// ReleaseAlibabaEleEnterpriseRestaurantGetAPIResponse 将 AlibabaEleEnterpriseRestaurantGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseRestaurantGetAPIResponse(v *AlibabaEleEnterpriseRestaurantGetAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseRestaurantGetAPIResponse.Put(v)
}
