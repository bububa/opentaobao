package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtItemGetAPIResponse 家装新零售商品信息查询 API返回值
// tmall.nrt.item.get
//
// 查询新零售商品信息
type TmallNrtItemGetAPIResponse struct {
	model.CommonResponse
	TmallNrtItemGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtItemGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtItemGetAPIResponseModel).Reset()
}

// TmallNrtItemGetAPIResponseModel is 家装新零售商品信息查询 成功返回结果
type TmallNrtItemGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	TmallNrtItemGet *TmallNrtItemGetResultDo `json:"tmall_nrt_item_get,omitempty" xml:"tmall_nrt_item_get,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtItemGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TmallNrtItemGet = nil
}

var poolTmallNrtItemGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtItemGetAPIResponse)
	},
}

// GetTmallNrtItemGetAPIResponse 从 sync.Pool 获取 TmallNrtItemGetAPIResponse
func GetTmallNrtItemGetAPIResponse() *TmallNrtItemGetAPIResponse {
	return poolTmallNrtItemGetAPIResponse.Get().(*TmallNrtItemGetAPIResponse)
}

// ReleaseTmallNrtItemGetAPIResponse 将 TmallNrtItemGetAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtItemGetAPIResponse(v *TmallNrtItemGetAPIResponse) {
	v.Reset()
	poolTmallNrtItemGetAPIResponse.Put(v)
}
