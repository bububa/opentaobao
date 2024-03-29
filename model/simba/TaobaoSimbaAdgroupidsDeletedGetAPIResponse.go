package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupidsDeletedGetAPIResponse 获取删除的推广组ID API返回值
// taobao.simba.adgroupids.deleted.get
//
// 获取删除的推广组ID
type TaobaoSimbaAdgroupidsDeletedGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAdgroupidsDeletedGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupidsDeletedGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaAdgroupidsDeletedGetAPIResponseModel).Reset()
}

// TaobaoSimbaAdgroupidsDeletedGetAPIResponseModel is 获取删除的推广组ID 成功返回结果
type TaobaoSimbaAdgroupidsDeletedGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroupids_deleted_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广组ID列表
	DeletedAdgroupIds []int64 `json:"deleted_adgroup_ids,omitempty" xml:"deleted_adgroup_ids>int64,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupidsDeletedGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeletedAdgroupIds = m.DeletedAdgroupIds[:0]
}

var poolTaobaoSimbaAdgroupidsDeletedGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaAdgroupidsDeletedGetAPIResponse)
	},
}

// GetTaobaoSimbaAdgroupidsDeletedGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaAdgroupidsDeletedGetAPIResponse
func GetTaobaoSimbaAdgroupidsDeletedGetAPIResponse() *TaobaoSimbaAdgroupidsDeletedGetAPIResponse {
	return poolTaobaoSimbaAdgroupidsDeletedGetAPIResponse.Get().(*TaobaoSimbaAdgroupidsDeletedGetAPIResponse)
}

// ReleaseTaobaoSimbaAdgroupidsDeletedGetAPIResponse 将 TaobaoSimbaAdgroupidsDeletedGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaAdgroupidsDeletedGetAPIResponse(v *TaobaoSimbaAdgroupidsDeletedGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaAdgroupidsDeletedGetAPIResponse.Put(v)
}
