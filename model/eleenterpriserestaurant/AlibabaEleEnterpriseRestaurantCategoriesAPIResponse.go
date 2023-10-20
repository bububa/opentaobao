package eleenterpriserestaurant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseRestaurantCategoriesAPIResponse 餐厅分类 API返回值
// alibaba.ele.enterprise.restaurant.categories
//
// 餐厅分类
type AlibabaEleEnterpriseRestaurantCategoriesAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseRestaurantCategoriesAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseRestaurantCategoriesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseRestaurantCategoriesAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseRestaurantCategoriesAPIResponseModel is 餐厅分类 成功返回结果
type AlibabaEleEnterpriseRestaurantCategoriesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_restaurant_categories_response"`
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
func (m *AlibabaEleEnterpriseRestaurantCategoriesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseDatas = m.EnterpriseDatas[:0]
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
}

var poolAlibabaEleEnterpriseRestaurantCategoriesAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseRestaurantCategoriesAPIResponse)
	},
}

// GetAlibabaEleEnterpriseRestaurantCategoriesAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseRestaurantCategoriesAPIResponse
func GetAlibabaEleEnterpriseRestaurantCategoriesAPIResponse() *AlibabaEleEnterpriseRestaurantCategoriesAPIResponse {
	return poolAlibabaEleEnterpriseRestaurantCategoriesAPIResponse.Get().(*AlibabaEleEnterpriseRestaurantCategoriesAPIResponse)
}

// ReleaseAlibabaEleEnterpriseRestaurantCategoriesAPIResponse 将 AlibabaEleEnterpriseRestaurantCategoriesAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseRestaurantCategoriesAPIResponse(v *AlibabaEleEnterpriseRestaurantCategoriesAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseRestaurantCategoriesAPIResponse.Put(v)
}
