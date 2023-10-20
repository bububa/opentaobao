package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupUpdateAPIResponse 更新一个推广组的信息 API返回值
// taobao.simba.adgroup.update
//
// 更新一个推广组的信息，可以设置默认出价、是否上线、非搜索出价、非搜索是否使用默认出价
type TaobaoSimbaAdgroupUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAdgroupUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaAdgroupUpdateAPIResponseModel).Reset()
}

// TaobaoSimbaAdgroupUpdateAPIResponseModel is 更新一个推广组的信息 成功返回结果
type TaobaoSimbaAdgroupUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroup_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 被修改的推广组
	Adgroup *ADGroup `json:"adgroup,omitempty" xml:"adgroup,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Adgroup = nil
}

var poolTaobaoSimbaAdgroupUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaAdgroupUpdateAPIResponse)
	},
}

// GetTaobaoSimbaAdgroupUpdateAPIResponse 从 sync.Pool 获取 TaobaoSimbaAdgroupUpdateAPIResponse
func GetTaobaoSimbaAdgroupUpdateAPIResponse() *TaobaoSimbaAdgroupUpdateAPIResponse {
	return poolTaobaoSimbaAdgroupUpdateAPIResponse.Get().(*TaobaoSimbaAdgroupUpdateAPIResponse)
}

// ReleaseTaobaoSimbaAdgroupUpdateAPIResponse 将 TaobaoSimbaAdgroupUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaAdgroupUpdateAPIResponse(v *TaobaoSimbaAdgroupUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSimbaAdgroupUpdateAPIResponse.Put(v)
}
