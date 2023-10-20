package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMixnickPlaytoweAPIResponse 互动mixNick转微淘 API返回值
// taobao.mixnick.playtowe
//
// 微淘应用的混淆nick转为互动类型混淆nick
type TaobaoMixnickPlaytoweAPIResponse struct {
	model.CommonResponse
	TaobaoMixnickPlaytoweAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMixnickPlaytoweAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMixnickPlaytoweAPIResponseModel).Reset()
}

// TaobaoMixnickPlaytoweAPIResponseModel is 互动mixNick转微淘 成功返回结果
type TaobaoMixnickPlaytoweAPIResponseModel struct {
	XMLName xml.Name `xml:"mixnick_playtowe_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 微淘混淆nick
	WeMixnick string `json:"we_mixnick,omitempty" xml:"we_mixnick,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMixnickPlaytoweAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WeMixnick = ""
}

var poolTaobaoMixnickPlaytoweAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMixnickPlaytoweAPIResponse)
	},
}

// GetTaobaoMixnickPlaytoweAPIResponse 从 sync.Pool 获取 TaobaoMixnickPlaytoweAPIResponse
func GetTaobaoMixnickPlaytoweAPIResponse() *TaobaoMixnickPlaytoweAPIResponse {
	return poolTaobaoMixnickPlaytoweAPIResponse.Get().(*TaobaoMixnickPlaytoweAPIResponse)
}

// ReleaseTaobaoMixnickPlaytoweAPIResponse 将 TaobaoMixnickPlaytoweAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMixnickPlaytoweAPIResponse(v *TaobaoMixnickPlaytoweAPIResponse) {
	v.Reset()
	poolTaobaoMixnickPlaytoweAPIResponse.Put(v)
}
