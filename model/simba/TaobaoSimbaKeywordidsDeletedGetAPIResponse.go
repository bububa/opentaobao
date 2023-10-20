package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordidsDeletedGetAPIResponse 获取删除的词ID API返回值
// taobao.simba.keywordids.deleted.get
//
// 获取删除的词ID
type TaobaoSimbaKeywordidsDeletedGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordidsDeletedGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordidsDeletedGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaKeywordidsDeletedGetAPIResponseModel).Reset()
}

// TaobaoSimbaKeywordidsDeletedGetAPIResponseModel is 获取删除的词ID 成功返回结果
type TaobaoSimbaKeywordidsDeletedGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywordids_deleted_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 词ID列表
	DeletedKeywordIds []int64 `json:"deleted_keyword_ids,omitempty" xml:"deleted_keyword_ids>int64,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordidsDeletedGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeletedKeywordIds = m.DeletedKeywordIds[:0]
}

var poolTaobaoSimbaKeywordidsDeletedGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordidsDeletedGetAPIResponse)
	},
}

// GetTaobaoSimbaKeywordidsDeletedGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaKeywordidsDeletedGetAPIResponse
func GetTaobaoSimbaKeywordidsDeletedGetAPIResponse() *TaobaoSimbaKeywordidsDeletedGetAPIResponse {
	return poolTaobaoSimbaKeywordidsDeletedGetAPIResponse.Get().(*TaobaoSimbaKeywordidsDeletedGetAPIResponse)
}

// ReleaseTaobaoSimbaKeywordidsDeletedGetAPIResponse 将 TaobaoSimbaKeywordidsDeletedGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaKeywordidsDeletedGetAPIResponse(v *TaobaoSimbaKeywordidsDeletedGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaKeywordidsDeletedGetAPIResponse.Put(v)
}
