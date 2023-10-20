package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupDeleteAPIResponse 删除一个推广组 API返回值
// taobao.simba.adgroup.delete
//
// 删除一个推广组
type TaobaoSimbaAdgroupDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAdgroupDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaAdgroupDeleteAPIResponseModel).Reset()
}

// TaobaoSimbaAdgroupDeleteAPIResponseModel is 删除一个推广组 成功返回结果
type TaobaoSimbaAdgroupDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroup_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 被删除的推广组
	Adgroup *ADGroup `json:"adgroup,omitempty" xml:"adgroup,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Adgroup = nil
}

var poolTaobaoSimbaAdgroupDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaAdgroupDeleteAPIResponse)
	},
}

// GetTaobaoSimbaAdgroupDeleteAPIResponse 从 sync.Pool 获取 TaobaoSimbaAdgroupDeleteAPIResponse
func GetTaobaoSimbaAdgroupDeleteAPIResponse() *TaobaoSimbaAdgroupDeleteAPIResponse {
	return poolTaobaoSimbaAdgroupDeleteAPIResponse.Get().(*TaobaoSimbaAdgroupDeleteAPIResponse)
}

// ReleaseTaobaoSimbaAdgroupDeleteAPIResponse 将 TaobaoSimbaAdgroupDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaAdgroupDeleteAPIResponse(v *TaobaoSimbaAdgroupDeleteAPIResponse) {
	v.Reset()
	poolTaobaoSimbaAdgroupDeleteAPIResponse.Put(v)
}
