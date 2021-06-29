package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
b2c订单状态驱动 API返回值 
alibaba.nlife.b2c.tradestatus.drive

用于驱动零售+订单状态
*/
type AlibabaNlifeB2cTradestatusDriveAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeB2cTradestatusDriveResponse
}

// b2c订单状态驱动 成功返回结果
type AlibabaNlifeB2cTradestatusDriveResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_b2c_tradestatus_drive_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 描述
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
