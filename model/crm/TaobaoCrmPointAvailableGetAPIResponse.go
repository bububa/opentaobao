package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmPointAvailableGetAPIResponse CRM会员积分查询开放接口 API返回值
// taobao.crm.point.available.get
//
// 查询用户在某个商家的可用积分数
type TaobaoCrmPointAvailableGetAPIResponse struct {
	model.CommonResponse
	TaobaoCrmPointAvailableGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmPointAvailableGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmPointAvailableGetAPIResponseModel).Reset()
}

// TaobaoCrmPointAvailableGetAPIResponseModel is CRM会员积分查询开放接口 成功返回结果
type TaobaoCrmPointAvailableGetAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_point_available_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 积分数
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmPointAvailableGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolTaobaoCrmPointAvailableGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmPointAvailableGetAPIResponse)
	},
}

// GetTaobaoCrmPointAvailableGetAPIResponse 从 sync.Pool 获取 TaobaoCrmPointAvailableGetAPIResponse
func GetTaobaoCrmPointAvailableGetAPIResponse() *TaobaoCrmPointAvailableGetAPIResponse {
	return poolTaobaoCrmPointAvailableGetAPIResponse.Get().(*TaobaoCrmPointAvailableGetAPIResponse)
}

// ReleaseTaobaoCrmPointAvailableGetAPIResponse 将 TaobaoCrmPointAvailableGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmPointAvailableGetAPIResponse(v *TaobaoCrmPointAvailableGetAPIResponse) {
	v.Reset()
	poolTaobaoCrmPointAvailableGetAPIResponse.Put(v)
}
