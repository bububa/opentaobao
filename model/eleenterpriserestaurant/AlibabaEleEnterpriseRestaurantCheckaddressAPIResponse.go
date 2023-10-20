package eleenterpriserestaurant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse 检查地址是否在餐厅配送范围内 API返回值
// alibaba.ele.enterprise.restaurant.checkaddress
//
// 检查地址是否在餐厅配送范围内
type AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseRestaurantCheckaddressAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseRestaurantCheckaddressAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseRestaurantCheckaddressAPIResponseModel is 检查地址是否在餐厅配送范围内 成功返回结果
type AlibabaEleEnterpriseRestaurantCheckaddressAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_restaurant_checkaddress_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值信息
	EnterpriseDatas []EnterpriseData `json:"enterprise_datas,omitempty" xml:"enterprise_datas>enterprise_data,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseRestaurantCheckaddressAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseDatas = m.EnterpriseDatas[:0]
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
}

var poolAlibabaEleEnterpriseRestaurantCheckaddressAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse)
	},
}

// GetAlibabaEleEnterpriseRestaurantCheckaddressAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse
func GetAlibabaEleEnterpriseRestaurantCheckaddressAPIResponse() *AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse {
	return poolAlibabaEleEnterpriseRestaurantCheckaddressAPIResponse.Get().(*AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse)
}

// ReleaseAlibabaEleEnterpriseRestaurantCheckaddressAPIResponse 将 AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseRestaurantCheckaddressAPIResponse(v *AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseRestaurantCheckaddressAPIResponse.Put(v)
}
