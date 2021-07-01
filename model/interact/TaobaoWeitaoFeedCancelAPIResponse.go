package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWeitaoFeedCancelAPIResponse
取消广播在timeline、广场中展示 API返回值
taobao.weitao.feed.cancel

取消广播在timeline和广场中的展示。 */
type TaobaoWeitaoFeedCancelAPIResponse struct {
	model.CommonResponse
	TaobaoWeitaoFeedCancelAPIResponseModel
}

// TaobaoWeitaoFeedCancelAPIResponseModel is 取消广播在timeline、广场中展示 成功返回结果
type TaobaoWeitaoFeedCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"weitao_feed_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推送结果
	Result *PushResult `json:"result,omitempty" xml:"result,omitempty"`
}
