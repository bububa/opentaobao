package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMixnickWetoplayAPIResponse 微淘混淆nick转互动混淆nick API返回值
// taobao.mixnick.wetoplay
//
// 微淘应用的混淆nick转为互动类型混淆nick
type TaobaoMixnickWetoplayAPIResponse struct {
	model.CommonResponse
	TaobaoMixnickWetoplayAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMixnickWetoplayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMixnickWetoplayAPIResponseModel).Reset()
}

// TaobaoMixnickWetoplayAPIResponseModel is 微淘混淆nick转互动混淆nick 成功返回结果
type TaobaoMixnickWetoplayAPIResponseModel struct {
	XMLName xml.Name `xml:"mixnick_wetoplay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 微淘转互动混淆nick
	PlayMixnickData *MixNickResult `json:"play_mixnick_data,omitempty" xml:"play_mixnick_data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMixnickWetoplayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PlayMixnickData = nil
}

var poolTaobaoMixnickWetoplayAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMixnickWetoplayAPIResponse)
	},
}

// GetTaobaoMixnickWetoplayAPIResponse 从 sync.Pool 获取 TaobaoMixnickWetoplayAPIResponse
func GetTaobaoMixnickWetoplayAPIResponse() *TaobaoMixnickWetoplayAPIResponse {
	return poolTaobaoMixnickWetoplayAPIResponse.Get().(*TaobaoMixnickWetoplayAPIResponse)
}

// ReleaseTaobaoMixnickWetoplayAPIResponse 将 TaobaoMixnickWetoplayAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMixnickWetoplayAPIResponse(v *TaobaoMixnickWetoplayAPIResponse) {
	v.Reset()
	poolTaobaoMixnickWetoplayAPIResponse.Put(v)
}
