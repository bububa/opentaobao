package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsMenuinfoReportAPIResponse 聚石塔菜单信息上报 API返回值
// taobao.jst.sms.menuinfo.report
//
// 聚石塔菜单信息上报
type TaobaoJstSmsMenuinfoReportAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsMenuinfoReportAPIResponseModel
}

// TaobaoJstSmsMenuinfoReportAPIResponseModel is 聚石塔菜单信息上报 成功返回结果
type TaobaoJstSmsMenuinfoReportAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_menuinfo_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统异常
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 请求id
	ResponseId string `json:"response_id,omitempty" xml:"response_id,omitempty"`
	// 系统异常
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
	// 成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
