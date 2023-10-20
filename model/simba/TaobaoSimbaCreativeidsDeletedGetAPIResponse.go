package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCreativeidsDeletedGetAPIResponse 获取删除的创意ID API返回值
// taobao.simba.creativeids.deleted.get
//
// 获取删除的创意ID
type TaobaoSimbaCreativeidsDeletedGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCreativeidsDeletedGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCreativeidsDeletedGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCreativeidsDeletedGetAPIResponseModel).Reset()
}

// TaobaoSimbaCreativeidsDeletedGetAPIResponseModel is 获取删除的创意ID 成功返回结果
type TaobaoSimbaCreativeidsDeletedGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_creativeids_deleted_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创意ID列表
	DeletedCreativeIds []int64 `json:"deleted_creative_ids,omitempty" xml:"deleted_creative_ids>int64,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCreativeidsDeletedGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeletedCreativeIds = m.DeletedCreativeIds[:0]
}

var poolTaobaoSimbaCreativeidsDeletedGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCreativeidsDeletedGetAPIResponse)
	},
}

// GetTaobaoSimbaCreativeidsDeletedGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaCreativeidsDeletedGetAPIResponse
func GetTaobaoSimbaCreativeidsDeletedGetAPIResponse() *TaobaoSimbaCreativeidsDeletedGetAPIResponse {
	return poolTaobaoSimbaCreativeidsDeletedGetAPIResponse.Get().(*TaobaoSimbaCreativeidsDeletedGetAPIResponse)
}

// ReleaseTaobaoSimbaCreativeidsDeletedGetAPIResponse 将 TaobaoSimbaCreativeidsDeletedGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCreativeidsDeletedGetAPIResponse(v *TaobaoSimbaCreativeidsDeletedGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCreativeidsDeletedGetAPIResponse.Put(v)
}
