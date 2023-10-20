package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimSnfilterwordSetfilterAPIResponse 关键词过滤 API返回值
// taobao.openim.snfilterword.setfilter
//
// 设置openim关键词过滤
type TaobaoOpenimSnfilterwordSetfilterAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimSnfilterwordSetfilterAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimSnfilterwordSetfilterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimSnfilterwordSetfilterAPIResponseModel).Reset()
}

// TaobaoOpenimSnfilterwordSetfilterAPIResponseModel is 关键词过滤 成功返回结果
type TaobaoOpenimSnfilterwordSetfilterAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_snfilterword_setfilter_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误原因
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 成功
	Errid int64 `json:"errid,omitempty" xml:"errid,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimSnfilterwordSetfilterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errmsg = ""
	m.Errid = 0
}

var poolTaobaoOpenimSnfilterwordSetfilterAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimSnfilterwordSetfilterAPIResponse)
	},
}

// GetTaobaoOpenimSnfilterwordSetfilterAPIResponse 从 sync.Pool 获取 TaobaoOpenimSnfilterwordSetfilterAPIResponse
func GetTaobaoOpenimSnfilterwordSetfilterAPIResponse() *TaobaoOpenimSnfilterwordSetfilterAPIResponse {
	return poolTaobaoOpenimSnfilterwordSetfilterAPIResponse.Get().(*TaobaoOpenimSnfilterwordSetfilterAPIResponse)
}

// ReleaseTaobaoOpenimSnfilterwordSetfilterAPIResponse 将 TaobaoOpenimSnfilterwordSetfilterAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimSnfilterwordSetfilterAPIResponse(v *TaobaoOpenimSnfilterwordSetfilterAPIResponse) {
	v.Reset()
	poolTaobaoOpenimSnfilterwordSetfilterAPIResponse.Put(v)
}
