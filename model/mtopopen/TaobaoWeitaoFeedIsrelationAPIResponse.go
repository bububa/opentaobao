package mtopopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeitaoFeedIsrelationAPIResponse 是否关注 API返回值
// taobao.weitao.feed.isrelation
//
// 判断用户是否关注对应的公共账号
type TaobaoWeitaoFeedIsrelationAPIResponse struct {
	model.CommonResponse
	TaobaoWeitaoFeedIsrelationAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWeitaoFeedIsrelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWeitaoFeedIsrelationAPIResponseModel).Reset()
}

// TaobaoWeitaoFeedIsrelationAPIResponseModel is 是否关注 成功返回结果
type TaobaoWeitaoFeedIsrelationAPIResponseModel struct {
	XMLName xml.Name `xml:"weitao_feed_isrelation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否关注
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWeitaoFeedIsrelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolTaobaoWeitaoFeedIsrelationAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWeitaoFeedIsrelationAPIResponse)
	},
}

// GetTaobaoWeitaoFeedIsrelationAPIResponse 从 sync.Pool 获取 TaobaoWeitaoFeedIsrelationAPIResponse
func GetTaobaoWeitaoFeedIsrelationAPIResponse() *TaobaoWeitaoFeedIsrelationAPIResponse {
	return poolTaobaoWeitaoFeedIsrelationAPIResponse.Get().(*TaobaoWeitaoFeedIsrelationAPIResponse)
}

// ReleaseTaobaoWeitaoFeedIsrelationAPIResponse 将 TaobaoWeitaoFeedIsrelationAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWeitaoFeedIsrelationAPIResponse(v *TaobaoWeitaoFeedIsrelationAPIResponse) {
	v.Reset()
	poolTaobaoWeitaoFeedIsrelationAPIResponse.Put(v)
}
