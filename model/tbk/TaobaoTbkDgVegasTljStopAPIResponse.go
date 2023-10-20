package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgvegastljstopAPIResponse 淘宝客-推广者-淘礼金暂停发放 API返回值
// taobao.tbk.dg.vegas.tlj.stop
//
// 淘宝客推广者可对已经创建的淘礼金暂停发放
type TaobaotbkdgvegastljstopAPIResponse struct {
	model.CommonResponse
	TaobaotbkdgvegastljstopAPIResponseModel
}

// TaobaotbkdgvegastljstopAPIResponseModel is 淘宝客-推广者-淘礼金暂停发放 成功返回结果
type TaobaotbkdgvegastljstopAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_vegas_tlj_stop_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgInfo
	Msginfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	Msgcode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model *UpdateStatusResult `json:"model,omitempty" xml:"model,omitempty"`
	// 调用接口是否成功
	Resultsuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
