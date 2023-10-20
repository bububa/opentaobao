package aliexpress

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressFeedPostPublishAPIResponse 同步帖子 API返回值
// aliexpress.feed.post.publish
//
// 站外平台同步帖子至AE FEED域
type AliexpressFeedPostPublishAPIResponse struct {
	model.CommonResponse
	AliexpressFeedPostPublishAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressFeedPostPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressFeedPostPublishAPIResponseModel).Reset()
}

// AliexpressFeedPostPublishAPIResponseModel is 同步帖子 成功返回结果
type AliexpressFeedPostPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_feed_post_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AliexpressFeedPostPublishResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressFeedPostPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressFeedPostPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressFeedPostPublishAPIResponse)
	},
}

// GetAliexpressFeedPostPublishAPIResponse 从 sync.Pool 获取 AliexpressFeedPostPublishAPIResponse
func GetAliexpressFeedPostPublishAPIResponse() *AliexpressFeedPostPublishAPIResponse {
	return poolAliexpressFeedPostPublishAPIResponse.Get().(*AliexpressFeedPostPublishAPIResponse)
}

// ReleaseAliexpressFeedPostPublishAPIResponse 将 AliexpressFeedPostPublishAPIResponse 保存到 sync.Pool
func ReleaseAliexpressFeedPostPublishAPIResponse(v *AliexpressFeedPostPublishAPIResponse) {
	v.Reset()
	poolAliexpressFeedPostPublishAPIResponse.Put(v)
}
