package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappMesssageReplyAPIResponse 轻店铺下行回复接口 API返回值
// taobao.miniapp.messsage.reply
//
// 外部 isv 调用该进口来进行轻店铺消息的回复
type TaobaoMiniappMesssageReplyAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappMesssageReplyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappMesssageReplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappMesssageReplyAPIResponseModel).Reset()
}

// TaobaoMiniappMesssageReplyAPIResponseModel is 轻店铺下行回复接口 成功返回结果
type TaobaoMiniappMesssageReplyAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_messsage_reply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoMiniappMesssageReplyResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappMesssageReplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoMiniappMesssageReplyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappMesssageReplyAPIResponse)
	},
}

// GetTaobaoMiniappMesssageReplyAPIResponse 从 sync.Pool 获取 TaobaoMiniappMesssageReplyAPIResponse
func GetTaobaoMiniappMesssageReplyAPIResponse() *TaobaoMiniappMesssageReplyAPIResponse {
	return poolTaobaoMiniappMesssageReplyAPIResponse.Get().(*TaobaoMiniappMesssageReplyAPIResponse)
}

// ReleaseTaobaoMiniappMesssageReplyAPIResponse 将 TaobaoMiniappMesssageReplyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappMesssageReplyAPIResponse(v *TaobaoMiniappMesssageReplyAPIResponse) {
	v.Reset()
	poolTaobaoMiniappMesssageReplyAPIResponse.Put(v)
}
