package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCustomersSidGetAPIResponse 查看功能权限 API返回值
// taobao.simba.customers.sid.get
//
// 查询用户是否拥有某个功能权限
type TaobaoSimbaCustomersSidGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCustomersSidGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCustomersSidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCustomersSidGetAPIResponseModel).Reset()
}

// TaobaoSimbaCustomersSidGetAPIResponseModel is 查看功能权限 成功返回结果
type TaobaoSimbaCustomersSidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_customers_sid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 权限列表及是否有权限
	Result *SidVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCustomersSidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSimbaCustomersSidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCustomersSidGetAPIResponse)
	},
}

// GetTaobaoSimbaCustomersSidGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaCustomersSidGetAPIResponse
func GetTaobaoSimbaCustomersSidGetAPIResponse() *TaobaoSimbaCustomersSidGetAPIResponse {
	return poolTaobaoSimbaCustomersSidGetAPIResponse.Get().(*TaobaoSimbaCustomersSidGetAPIResponse)
}

// ReleaseTaobaoSimbaCustomersSidGetAPIResponse 将 TaobaoSimbaCustomersSidGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCustomersSidGetAPIResponse(v *TaobaoSimbaCustomersSidGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCustomersSidGetAPIResponse.Put(v)
}
