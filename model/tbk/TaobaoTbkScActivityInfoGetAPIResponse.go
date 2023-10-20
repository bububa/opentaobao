package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScActivityInfoGetAPIResponse 淘宝客-服务商-官方活动转链 API返回值
// taobao.tbk.sc.activity.info.get
//
// 服务商使用。支持入参推广者对应的推广位和官方活动会场ID，获取活动信息和推广者的推广链接，包含推广长链接、短链接、淘口令、微信推广二维码地址等。改该API支持二方、三方类型的转链。官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取。
type TaobaoTbkScActivityInfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScActivityInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScActivityInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScActivityInfoGetAPIResponseModel).Reset()
}

// TaobaoTbkScActivityInfoGetAPIResponseModel is 淘宝客-服务商-官方活动转链 成功返回结果
type TaobaoTbkScActivityInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_activity_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Data *TaobaoTbkScActivityInfoGetData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScActivityInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkScActivityInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScActivityInfoGetAPIResponse)
	},
}

// GetTaobaoTbkScActivityInfoGetAPIResponse 从 sync.Pool 获取 TaobaoTbkScActivityInfoGetAPIResponse
func GetTaobaoTbkScActivityInfoGetAPIResponse() *TaobaoTbkScActivityInfoGetAPIResponse {
	return poolTaobaoTbkScActivityInfoGetAPIResponse.Get().(*TaobaoTbkScActivityInfoGetAPIResponse)
}

// ReleaseTaobaoTbkScActivityInfoGetAPIResponse 将 TaobaoTbkScActivityInfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScActivityInfoGetAPIResponse(v *TaobaoTbkScActivityInfoGetAPIResponse) {
	v.Reset()
	poolTaobaoTbkScActivityInfoGetAPIResponse.Put(v)
}
