package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappMesssageNormalSendAPIResponse 轻店铺下行普通消息给用户 API返回值
// taobao.miniapp.messsage.normal.send
//
// 小程序下行单个普通消息
type TaobaoMiniappMesssageNormalSendAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappMesssageNormalSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappMesssageNormalSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappMesssageNormalSendAPIResponseModel).Reset()
}

// TaobaoMiniappMesssageNormalSendAPIResponseModel is 轻店铺下行普通消息给用户 成功返回结果
type TaobaoMiniappMesssageNormalSendAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_messsage_normal_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回消息Id
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappMesssageNormalSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
}

var poolTaobaoMiniappMesssageNormalSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappMesssageNormalSendAPIResponse)
	},
}

// GetTaobaoMiniappMesssageNormalSendAPIResponse 从 sync.Pool 获取 TaobaoMiniappMesssageNormalSendAPIResponse
func GetTaobaoMiniappMesssageNormalSendAPIResponse() *TaobaoMiniappMesssageNormalSendAPIResponse {
	return poolTaobaoMiniappMesssageNormalSendAPIResponse.Get().(*TaobaoMiniappMesssageNormalSendAPIResponse)
}

// ReleaseTaobaoMiniappMesssageNormalSendAPIResponse 将 TaobaoMiniappMesssageNormalSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappMesssageNormalSendAPIResponse(v *TaobaoMiniappMesssageNormalSendAPIResponse) {
	v.Reset()
	poolTaobaoMiniappMesssageNormalSendAPIResponse.Put(v)
}
