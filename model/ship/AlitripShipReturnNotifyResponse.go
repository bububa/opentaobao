package ship

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
船票退票退款回填接口 API返回值 
alitrip.ship.return.notify

此接口为接入商调用飞猪接口回填退票状态，飞猪平台给用户进行退票退款。飞猪平台保证数据幂等。
*/
type AlitripShipReturnNotifyAPIResponse struct {
    model.CommonResponse
    AlitripShipReturnNotifyResponse
}

// 船票退票退款回填接口 成功返回结果
type AlitripShipReturnNotifyResponse struct {
    XMLName xml.Name `xml:"alitrip_ship_return_notify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误码
    RetCode   string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
    // 错误描述
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
    // 结果
    RetSuccess   bool `json:"ret_success,omitempty" xml:"ret_success,omitempty"`
}
