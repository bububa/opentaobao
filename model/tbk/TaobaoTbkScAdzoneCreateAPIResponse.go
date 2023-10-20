package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScAdzoneCreateAPIResponse 淘宝客-服务商-创建推广者位 API返回值
// taobao.tbk.sc.adzone.create
//
// 提供淘宝客创建广告位
type TaobaoTbkScAdzoneCreateAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScAdzoneCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScAdzoneCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScAdzoneCreateAPIResponseModel).Reset()
}

// TaobaoTbkScAdzoneCreateAPIResponseModel is 淘宝客-服务商-创建推广者位 成功返回结果
type TaobaoTbkScAdzoneCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_adzone_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// MapData
	Data *TaobaoTbkScAdzoneCreateMapData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScAdzoneCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkScAdzoneCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScAdzoneCreateAPIResponse)
	},
}

// GetTaobaoTbkScAdzoneCreateAPIResponse 从 sync.Pool 获取 TaobaoTbkScAdzoneCreateAPIResponse
func GetTaobaoTbkScAdzoneCreateAPIResponse() *TaobaoTbkScAdzoneCreateAPIResponse {
	return poolTaobaoTbkScAdzoneCreateAPIResponse.Get().(*TaobaoTbkScAdzoneCreateAPIResponse)
}

// ReleaseTaobaoTbkScAdzoneCreateAPIResponse 将 TaobaoTbkScAdzoneCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScAdzoneCreateAPIResponse(v *TaobaoTbkScAdzoneCreateAPIResponse) {
	v.Reset()
	poolTaobaoTbkScAdzoneCreateAPIResponse.Put(v)
}
