package tbk

import (
	"encoding/xml"

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

// TaobaoTbkScAdzoneCreateAPIResponseModel is 淘宝客-服务商-创建推广者位 成功返回结果
type TaobaoTbkScAdzoneCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_adzone_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// MapData
	Data *TaobaoTbkScAdzoneCreateMapData `json:"data,omitempty" xml:"data,omitempty"`
}
