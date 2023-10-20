package eleenterpriserestaurant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse 查询餐厅菜单 API返回值
// alibaba.ele.enterprise.restaurant.multispecmenu
//
// 查询餐厅菜单
type AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponseModel is 查询餐厅菜单 成功返回结果
type AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_restaurant_multispecmenu_response"`
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
func (m *AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseDatas = m.EnterpriseDatas[:0]
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
}

var poolAlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse)
	},
}

// GetAlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse
func GetAlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse() *AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse {
	return poolAlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse.Get().(*AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse)
}

// ReleaseAlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse 将 AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse(v *AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse.Put(v)
}
