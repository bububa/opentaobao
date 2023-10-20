package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeitaoFeedCancelAPIResponse 取消广播在timeline、广场中展示 API返回值
// taobao.weitao.feed.cancel
//
// 取消广播在timeline和广场中的展示。
type TaobaoWeitaoFeedCancelAPIResponse struct {
	model.CommonResponse
	TaobaoWeitaoFeedCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWeitaoFeedCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWeitaoFeedCancelAPIResponseModel).Reset()
}

// TaobaoWeitaoFeedCancelAPIResponseModel is 取消广播在timeline、广场中展示 成功返回结果
type TaobaoWeitaoFeedCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"weitao_feed_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推送结果
	Result *PushResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWeitaoFeedCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWeitaoFeedCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWeitaoFeedCancelAPIResponse)
	},
}

// GetTaobaoWeitaoFeedCancelAPIResponse 从 sync.Pool 获取 TaobaoWeitaoFeedCancelAPIResponse
func GetTaobaoWeitaoFeedCancelAPIResponse() *TaobaoWeitaoFeedCancelAPIResponse {
	return poolTaobaoWeitaoFeedCancelAPIResponse.Get().(*TaobaoWeitaoFeedCancelAPIResponse)
}

// ReleaseTaobaoWeitaoFeedCancelAPIResponse 将 TaobaoWeitaoFeedCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWeitaoFeedCancelAPIResponse(v *TaobaoWeitaoFeedCancelAPIResponse) {
	v.Reset()
	poolTaobaoWeitaoFeedCancelAPIResponse.Put(v)
}
