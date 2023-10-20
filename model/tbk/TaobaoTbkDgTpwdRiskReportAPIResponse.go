package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgtpwdriskreportAPIResponse 淘宝客-推广者-淘口令预警及拦截查询 API返回值
// taobao.tbk.dg.tpwd.risk.report
//
// 淘宝客-推广者-淘口令预警及拦截查询
type TaobaotbkdgtpwdriskreportAPIResponse struct {
	model.CommonResponse
	TaobaotbkdgtpwdriskreportAPIResponseModel
}

// TaobaotbkdgtpwdriskreportAPIResponseModel is 淘宝客-推广者-淘口令预警及拦截查询 成功返回结果
type TaobaotbkdgtpwdriskreportAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_tpwd_risk_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无
	Result *TaobaotbkdgtpwdriskreportResult `json:"result,omitempty" xml:"result,omitempty"`
}
