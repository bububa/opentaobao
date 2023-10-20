package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkActivityInfoGetAPIResponse 淘宝客-推广者-官方活动转链 API返回值
// taobao.tbk.activity.info.get
//
// 支持入参推广位和官方活动会场ID，获取活动信息和推广链接，包含推广长链接、短链接、淘口令、微信推广二维码地址等。改该API支持二方、三方类型的转链。官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取。
type TaobaoTbkActivityInfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkActivityInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkActivityInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkActivityInfoGetAPIResponseModel).Reset()
}

// TaobaoTbkActivityInfoGetAPIResponseModel is 淘宝客-推广者-官方活动转链 成功返回结果
type TaobaoTbkActivityInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_activity_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Data *TaobaoTbkActivityInfoGetData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkActivityInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkActivityInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkActivityInfoGetAPIResponse)
	},
}

// GetTaobaoTbkActivityInfoGetAPIResponse 从 sync.Pool 获取 TaobaoTbkActivityInfoGetAPIResponse
func GetTaobaoTbkActivityInfoGetAPIResponse() *TaobaoTbkActivityInfoGetAPIResponse {
	return poolTaobaoTbkActivityInfoGetAPIResponse.Get().(*TaobaoTbkActivityInfoGetAPIResponse)
}

// ReleaseTaobaoTbkActivityInfoGetAPIResponse 将 TaobaoTbkActivityInfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkActivityInfoGetAPIResponse(v *TaobaoTbkActivityInfoGetAPIResponse) {
	v.Reset()
	poolTaobaoTbkActivityInfoGetAPIResponse.Put(v)
}
