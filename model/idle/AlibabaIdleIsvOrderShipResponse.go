package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼无忧购服务商发货 APIResponse
alibaba.idle.isv.order.ship

闲鱼无忧购业务入仓模式下服务商订单发货的接口
*/
type AlibabaIdleIsvOrderShipAPIResponse struct {
    model.CommonResponse
    AlibabaIdleIsvOrderShipResponse
}

type AlibabaIdleIsvOrderShipResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_isv_order_ship_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求结果，是否成功
    
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`

    
    // 错误码
    
    ResultErrCode   string `json:"result_err_code,omitempty" xml:"result_err_code,omitempty"`

    
    // 业务逻辑结果，暂时不用
    
    ResultModule   bool `json:"result_module,omitempty" xml:"result_module,omitempty"`

    
    // 错误信息
    
    ResultErrMsg   string `json:"result_err_msg,omitempty" xml:"result_err_msg,omitempty"`

    
}
