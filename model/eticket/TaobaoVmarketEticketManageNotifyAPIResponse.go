package eticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaovmarketeticketmanagenotifyAPIResponse 主动发起通知接口 API返回值
// taobao.vmarket.eticket.manage.notify
//
// 外部合作商家主动发起通知接口
type TaobaovmarketeticketmanagenotifyAPIResponse struct {
	model.CommonResponse
	TaobaovmarketeticketmanagenotifyAPIResponseModel
}

// TaobaovmarketeticketmanagenotifyAPIResponseModel is 主动发起通知接口 成功返回结果
type TaobaovmarketeticketmanagenotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_manage_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1:成功
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}
