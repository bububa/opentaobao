package nlife

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// Alibabanlifeb2ctradestatusdriveAPIResponse b2c订单状态驱动 API返回值
// alibaba.nlife.b2c.tradestatus.drive
//
// 用于驱动零售+订单状态
type Alibabanlifeb2ctradestatusdriveAPIResponse struct {
	model.CommonResponse
	Alibabanlifeb2ctradestatusdriveAPIResponseModel
}

// Alibabanlifeb2ctradestatusdriveAPIResponseModel is b2c订单状态驱动 成功返回结果
type Alibabanlifeb2ctradestatusdriveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_tradestatus_drive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
