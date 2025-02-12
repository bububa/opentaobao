package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgGeneralLinkConvertAPIResponse 淘宝客-推广者-万能转链 API返回值
// taobao.tbk.dg.general.link.convert
//
// 淘宝客-推广者-万能转链
type TaobaoTbkDgGeneralLinkConvertAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgGeneralLinkConvertAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkDgGeneralLinkConvertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkDgGeneralLinkConvertAPIResponseModel).Reset()
}

// TaobaoTbkDgGeneralLinkConvertAPIResponseModel is 淘宝客-推广者-万能转链 成功返回结果
type TaobaoTbkDgGeneralLinkConvertAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_general_link_convert_response"`
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
func (m *TaobaoTbkDgGeneralLinkConvertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.Data = nil
	m.BizErrorDesc = 0
}

var poolTaobaoTbkDgGeneralLinkConvertAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgGeneralLinkConvertAPIResponse)
	},
}

// GetTaobaoTbkDgGeneralLinkConvertAPIResponse 从 sync.Pool 获取 TaobaoTbkDgGeneralLinkConvertAPIResponse
func GetTaobaoTbkDgGeneralLinkConvertAPIResponse() *TaobaoTbkDgGeneralLinkConvertAPIResponse {
	return poolTaobaoTbkDgGeneralLinkConvertAPIResponse.Get().(*TaobaoTbkDgGeneralLinkConvertAPIResponse)
}

// ReleaseTaobaoTbkDgGeneralLinkConvertAPIResponse 将 TaobaoTbkDgGeneralLinkConvertAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkDgGeneralLinkConvertAPIResponse(v *TaobaoTbkDgGeneralLinkConvertAPIResponse) {
	v.Reset()
	poolTaobaoTbkDgGeneralLinkConvertAPIResponse.Put(v)
}
