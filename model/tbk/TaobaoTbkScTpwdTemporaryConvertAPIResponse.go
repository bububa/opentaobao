package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScTpwdTemporaryConvertAPIResponse 淘宝客-服务商-淘口令解析&&转链（临时接口） API返回值
// taobao.tbk.sc.tpwd.temporary.convert
//
// 支持通过淘口令解析商品id，并提供对应的淘客转链接
type TaobaoTbkScTpwdTemporaryConvertAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScTpwdTemporaryConvertAPIResponseModel
}

// TaobaoTbkScTpwdTemporaryConvertAPIResponseModel is 淘宝客-服务商-淘口令解析&&转链（临时接口） 成功返回结果
type TaobaoTbkScTpwdTemporaryConvertAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_tpwd_temporary_convert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
