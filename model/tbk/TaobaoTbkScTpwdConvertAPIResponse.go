package tbk

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoTbkScTpwdConvertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScTpwdConvertAPIResponseModel).Reset()
}

// TaobaoTbkScTpwdConvertAPIResponseModel is 淘宝客-服务商-淘口令解析&转链 成功返回结果
type TaobaoTbkScTpwdConvertAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_tpwd_convert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaoTbkScTpwdConvertMapData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScTpwdConvertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkScTpwdConvertAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScTpwdConvertAPIResponse)
	},
}

// GetTaobaoTbkScTpwdConvertAPIResponse 从 sync.Pool 获取 TaobaoTbkScTpwdConvertAPIResponse
func GetTaobaoTbkScTpwdConvertAPIResponse() *TaobaoTbkScTpwdConvertAPIResponse {
	return poolTaobaoTbkScTpwdConvertAPIResponse.Get().(*TaobaoTbkScTpwdConvertAPIResponse)
}

// ReleaseTaobaoTbkScTpwdConvertAPIResponse 将 TaobaoTbkScTpwdConvertAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScTpwdConvertAPIResponse(v *TaobaoTbkScTpwdConvertAPIResponse) {
	v.Reset()
	poolTaobaoTbkScTpwdConvertAPIResponse.Put(v)
}
