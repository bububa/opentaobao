package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscvegassendstatusAPIResponse 淘宝客-服务商-红包领取状态查询 API返回值
// taobao.tbk.sc.vegas.send.status
//
// 服务商使用。支持淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包。接入前需签署协议：https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin
type TaobaotbkscvegassendstatusAPIResponse struct {
	model.CommonResponse
	TaobaotbkscvegassendstatusAPIResponseModel
}

// TaobaotbkscvegassendstatusAPIResponseModel is 淘宝客-服务商-红包领取状态查询 成功返回结果
type TaobaotbkscvegassendstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_vegas_send_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果描述信息
	Resultmsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回结果封装对象
	Data *TaobaotbkscvegassendstatusData `json:"data,omitempty" xml:"data,omitempty"`
}
