package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgGeneralLinkParseAPIResponse 淘宝客-推广者-万能解析 API返回值
// taobao.tbk.dg.general.link.parse
//
// 淘宝客-推广者-万能转链
type TaobaoTbkDgGeneralLinkParseAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgGeneralLinkParseAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkDgGeneralLinkParseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkDgGeneralLinkParseAPIResponseModel).Reset()
}

// TaobaoTbkDgGeneralLinkParseAPIResponseModel is 淘宝客-推广者-万能解析 成功返回结果
type TaobaoTbkDgGeneralLinkParseAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_general_link_parse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 见错误码描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// TbkLinkDTO
	Data *TbkLinkDto `json:"data,omitempty" xml:"data,omitempty"`
	// 见错误码描述
	BizErrorDesc int64 `json:"biz_error_desc,omitempty" xml:"biz_error_desc,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkDgGeneralLinkParseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.Data = nil
	m.BizErrorDesc = 0
}

var poolTaobaoTbkDgGeneralLinkParseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgGeneralLinkParseAPIResponse)
	},
}

// GetTaobaoTbkDgGeneralLinkParseAPIResponse 从 sync.Pool 获取 TaobaoTbkDgGeneralLinkParseAPIResponse
func GetTaobaoTbkDgGeneralLinkParseAPIResponse() *TaobaoTbkDgGeneralLinkParseAPIResponse {
	return poolTaobaoTbkDgGeneralLinkParseAPIResponse.Get().(*TaobaoTbkDgGeneralLinkParseAPIResponse)
}

// ReleaseTaobaoTbkDgGeneralLinkParseAPIResponse 将 TaobaoTbkDgGeneralLinkParseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkDgGeneralLinkParseAPIResponse(v *TaobaoTbkDgGeneralLinkParseAPIResponse) {
	v.Reset()
	poolTaobaoTbkDgGeneralLinkParseAPIResponse.Put(v)
}
