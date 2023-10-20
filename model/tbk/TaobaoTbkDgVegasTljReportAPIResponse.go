package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgVegasTljReportAPIResponse 淘宝客-推广者-淘礼金效果数据 API返回值
// taobao.tbk.dg.vegas.tlj.report
//
// 淘宝客推广者可查询淘礼金发放和使用等效果数据，只提供近150天的数据
type TaobaoTbkDgVegasTljReportAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgVegasTljReportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkDgVegasTljReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkDgVegasTljReportAPIResponseModel).Reset()
}

// TaobaoTbkDgVegasTljReportAPIResponseModel is 淘宝客-推广者-淘礼金效果数据 成功返回结果
type TaobaoTbkDgVegasTljReportAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_vegas_tlj_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 结果
	Model *InstanceDto `json:"model,omitempty" xml:"model,omitempty"`
	// 调用接口是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkDgVegasTljReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Model = nil
	m.ResultSuccess = false
}

var poolTaobaoTbkDgVegasTljReportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgVegasTljReportAPIResponse)
	},
}

// GetTaobaoTbkDgVegasTljReportAPIResponse 从 sync.Pool 获取 TaobaoTbkDgVegasTljReportAPIResponse
func GetTaobaoTbkDgVegasTljReportAPIResponse() *TaobaoTbkDgVegasTljReportAPIResponse {
	return poolTaobaoTbkDgVegasTljReportAPIResponse.Get().(*TaobaoTbkDgVegasTljReportAPIResponse)
}

// ReleaseTaobaoTbkDgVegasTljReportAPIResponse 将 TaobaoTbkDgVegasTljReportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkDgVegasTljReportAPIResponse(v *TaobaoTbkDgVegasTljReportAPIResponse) {
	v.Reset()
	poolTaobaoTbkDgVegasTljReportAPIResponse.Put(v)
}
