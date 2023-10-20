package shop

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoStoreFollowurlGetAPIResponse 获取店铺关注URL API返回值
// taobao.store.followurl.get
//
// 获取关注店铺的URL
type TaobaoStoreFollowurlGetAPIResponse struct {
	model.CommonResponse
	TaobaoStoreFollowurlGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoStoreFollowurlGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoStoreFollowurlGetAPIResponseModel).Reset()
}

// TaobaoStoreFollowurlGetAPIResponseModel is 获取店铺关注URL 成功返回结果
type TaobaoStoreFollowurlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"store_followurl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 店铺关注URL
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoStoreFollowurlGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Url = ""
}

var poolTaobaoStoreFollowurlGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoStoreFollowurlGetAPIResponse)
	},
}

// GetTaobaoStoreFollowurlGetAPIResponse 从 sync.Pool 获取 TaobaoStoreFollowurlGetAPIResponse
func GetTaobaoStoreFollowurlGetAPIResponse() *TaobaoStoreFollowurlGetAPIResponse {
	return poolTaobaoStoreFollowurlGetAPIResponse.Get().(*TaobaoStoreFollowurlGetAPIResponse)
}

// ReleaseTaobaoStoreFollowurlGetAPIResponse 将 TaobaoStoreFollowurlGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoStoreFollowurlGetAPIResponse(v *TaobaoStoreFollowurlGetAPIResponse) {
	v.Reset()
	poolTaobaoStoreFollowurlGetAPIResponse.Put(v)
}
