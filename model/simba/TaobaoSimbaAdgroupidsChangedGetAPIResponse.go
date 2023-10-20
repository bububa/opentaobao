package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupidsChangedGetAPIResponse 获取修改的推广组ID API返回值
// taobao.simba.adgroupids.changed.get
//
// 获取修改的推广组ID
type TaobaoSimbaAdgroupidsChangedGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAdgroupidsChangedGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupidsChangedGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaAdgroupidsChangedGetAPIResponseModel).Reset()
}

// TaobaoSimbaAdgroupidsChangedGetAPIResponseModel is 获取修改的推广组ID 成功返回结果
type TaobaoSimbaAdgroupidsChangedGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroupids_changed_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广组ID列表
	ChangedAdgroupids []int64 `json:"changed_adgroupids,omitempty" xml:"changed_adgroupids>int64,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupidsChangedGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ChangedAdgroupids = m.ChangedAdgroupids[:0]
}

var poolTaobaoSimbaAdgroupidsChangedGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaAdgroupidsChangedGetAPIResponse)
	},
}

// GetTaobaoSimbaAdgroupidsChangedGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaAdgroupidsChangedGetAPIResponse
func GetTaobaoSimbaAdgroupidsChangedGetAPIResponse() *TaobaoSimbaAdgroupidsChangedGetAPIResponse {
	return poolTaobaoSimbaAdgroupidsChangedGetAPIResponse.Get().(*TaobaoSimbaAdgroupidsChangedGetAPIResponse)
}

// ReleaseTaobaoSimbaAdgroupidsChangedGetAPIResponse 将 TaobaoSimbaAdgroupidsChangedGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaAdgroupidsChangedGetAPIResponse(v *TaobaoSimbaAdgroupidsChangedGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaAdgroupidsChangedGetAPIResponse.Put(v)
}
