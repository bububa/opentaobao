package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgvegastljreportAPIResponse 淘宝客-推广者-淘礼金效果数据 API返回值
// taobao.tbk.dg.vegas.tlj.report
//
// 淘宝客推广者可查询淘礼金发放和使用等效果数据，只提供近150天的数据
type TaobaotbkdgvegastljreportAPIResponse struct {
	model.CommonResponse
	TaobaotbkdgvegastljreportAPIResponseModel
}

// TaobaotbkdgvegastljreportAPIResponseModel is 淘宝客-推广者-淘礼金效果数据 成功返回结果
type TaobaotbkdgvegastljreportAPIResponseModel struct {
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
