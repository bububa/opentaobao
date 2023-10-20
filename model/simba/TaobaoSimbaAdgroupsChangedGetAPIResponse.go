package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupsChangedGetAPIResponse 分页获取修改的推广组ID和修改时间 API返回值
// taobao.simba.adgroups.changed.get
//
// 分页获取修改的推广组ID和修改时间
type TaobaoSimbaAdgroupsChangedGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAdgroupsChangedGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupsChangedGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaAdgroupsChangedGetAPIResponseModel).Reset()
}

// TaobaoSimbaAdgroupsChangedGetAPIResponseModel is 分页获取修改的推广组ID和修改时间 成功返回结果
type TaobaoSimbaAdgroupsChangedGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroups_changed_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广组分页对象
	Adgroups *ADGroupPage `json:"adgroups,omitempty" xml:"adgroups,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupsChangedGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Adgroups = nil
}

var poolTaobaoSimbaAdgroupsChangedGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaAdgroupsChangedGetAPIResponse)
	},
}

// GetTaobaoSimbaAdgroupsChangedGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaAdgroupsChangedGetAPIResponse
func GetTaobaoSimbaAdgroupsChangedGetAPIResponse() *TaobaoSimbaAdgroupsChangedGetAPIResponse {
	return poolTaobaoSimbaAdgroupsChangedGetAPIResponse.Get().(*TaobaoSimbaAdgroupsChangedGetAPIResponse)
}

// ReleaseTaobaoSimbaAdgroupsChangedGetAPIResponse 将 TaobaoSimbaAdgroupsChangedGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaAdgroupsChangedGetAPIResponse(v *TaobaoSimbaAdgroupsChangedGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaAdgroupsChangedGetAPIResponse.Put(v)
}
