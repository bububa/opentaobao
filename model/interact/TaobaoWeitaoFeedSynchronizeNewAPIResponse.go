package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeitaoFeedSynchronizeNewAPIResponse 推广淘小铺isv 活动到微淘feed API返回值
// taobao.weitao.feed.synchronize.new
//
// 推广微淘互动应用活动到微淘
type TaobaoWeitaoFeedSynchronizeNewAPIResponse struct {
	model.CommonResponse
	TaobaoWeitaoFeedSynchronizeNewAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWeitaoFeedSynchronizeNewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWeitaoFeedSynchronizeNewAPIResponseModel).Reset()
}

// TaobaoWeitaoFeedSynchronizeNewAPIResponseModel is 推广淘小铺isv 活动到微淘feed 成功返回结果
type TaobaoWeitaoFeedSynchronizeNewAPIResponseModel struct {
	XMLName xml.Name `xml:"weitao_feed_synchronize_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 增加错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 推送结果
	Result *PushResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWeitaoFeedSynchronizeNewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errmsg = ""
	m.Result = nil
}

var poolTaobaoWeitaoFeedSynchronizeNewAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWeitaoFeedSynchronizeNewAPIResponse)
	},
}

// GetTaobaoWeitaoFeedSynchronizeNewAPIResponse 从 sync.Pool 获取 TaobaoWeitaoFeedSynchronizeNewAPIResponse
func GetTaobaoWeitaoFeedSynchronizeNewAPIResponse() *TaobaoWeitaoFeedSynchronizeNewAPIResponse {
	return poolTaobaoWeitaoFeedSynchronizeNewAPIResponse.Get().(*TaobaoWeitaoFeedSynchronizeNewAPIResponse)
}

// ReleaseTaobaoWeitaoFeedSynchronizeNewAPIResponse 将 TaobaoWeitaoFeedSynchronizeNewAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWeitaoFeedSynchronizeNewAPIResponse(v *TaobaoWeitaoFeedSynchronizeNewAPIResponse) {
	v.Reset()
	poolTaobaoWeitaoFeedSynchronizeNewAPIResponse.Put(v)
}
