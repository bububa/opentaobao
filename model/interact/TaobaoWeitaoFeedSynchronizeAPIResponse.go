package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeitaoFeedSynchronizeAPIResponse 推广淘小铺isv 活动到微淘feed API返回值
// taobao.weitao.feed.synchronize
//
// 推广淘小铺isv 活动到微淘feed
type TaobaoWeitaoFeedSynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoWeitaoFeedSynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWeitaoFeedSynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWeitaoFeedSynchronizeAPIResponseModel).Reset()
}

// TaobaoWeitaoFeedSynchronizeAPIResponseModel is 推广淘小铺isv 活动到微淘feed 成功返回结果
type TaobaoWeitaoFeedSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"weitao_feed_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 增加错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 同步到微淘成功与否
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWeitaoFeedSynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errmsg = ""
	m.Result = false
}

var poolTaobaoWeitaoFeedSynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWeitaoFeedSynchronizeAPIResponse)
	},
}

// GetTaobaoWeitaoFeedSynchronizeAPIResponse 从 sync.Pool 获取 TaobaoWeitaoFeedSynchronizeAPIResponse
func GetTaobaoWeitaoFeedSynchronizeAPIResponse() *TaobaoWeitaoFeedSynchronizeAPIResponse {
	return poolTaobaoWeitaoFeedSynchronizeAPIResponse.Get().(*TaobaoWeitaoFeedSynchronizeAPIResponse)
}

// ReleaseTaobaoWeitaoFeedSynchronizeAPIResponse 将 TaobaoWeitaoFeedSynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWeitaoFeedSynchronizeAPIResponse(v *TaobaoWeitaoFeedSynchronizeAPIResponse) {
	v.Reset()
	poolTaobaoWeitaoFeedSynchronizeAPIResponse.Put(v)
}
