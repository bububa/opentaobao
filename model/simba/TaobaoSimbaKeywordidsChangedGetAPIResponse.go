package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordidsChangedGetAPIResponse 获取修改的词ID API返回值
// taobao.simba.keywordids.changed.get
//
// 获取修改的词ID
type TaobaoSimbaKeywordidsChangedGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordidsChangedGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordidsChangedGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaKeywordidsChangedGetAPIResponseModel).Reset()
}

// TaobaoSimbaKeywordidsChangedGetAPIResponseModel is 获取修改的词ID 成功返回结果
type TaobaoSimbaKeywordidsChangedGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywordids_changed_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 词的ID列表
	ChangedKeywordIds []int64 `json:"changed_keyword_ids,omitempty" xml:"changed_keyword_ids>int64,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordidsChangedGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ChangedKeywordIds = m.ChangedKeywordIds[:0]
}

var poolTaobaoSimbaKeywordidsChangedGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordidsChangedGetAPIResponse)
	},
}

// GetTaobaoSimbaKeywordidsChangedGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaKeywordidsChangedGetAPIResponse
func GetTaobaoSimbaKeywordidsChangedGetAPIResponse() *TaobaoSimbaKeywordidsChangedGetAPIResponse {
	return poolTaobaoSimbaKeywordidsChangedGetAPIResponse.Get().(*TaobaoSimbaKeywordidsChangedGetAPIResponse)
}

// ReleaseTaobaoSimbaKeywordidsChangedGetAPIResponse 将 TaobaoSimbaKeywordidsChangedGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaKeywordidsChangedGetAPIResponse(v *TaobaoSimbaKeywordidsChangedGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaKeywordidsChangedGetAPIResponse.Put(v)
}
