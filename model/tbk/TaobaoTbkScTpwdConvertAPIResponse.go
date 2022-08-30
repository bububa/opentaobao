package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScTpwdConvertAPIResponse 淘宝客-服务商-淘口令解析&转链 API返回值
// taobao.tbk.sc.tpwd.convert
//
// 支持通过淘口令解析商品id，并提供对应的淘客转链接
type TaobaoTbkScTpwdConvertAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScTpwdConvertAPIResponseModel
}

// TaobaoTbkScTpwdConvertAPIResponseModel is 淘宝客-服务商-淘口令解析&转链 成功返回结果
type TaobaoTbkScTpwdConvertAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_tpwd_convert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaoTbkScTpwdConvertMapData `json:"data,omitempty" xml:"data,omitempty"`
}
