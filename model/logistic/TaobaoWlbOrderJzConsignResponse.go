package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
家装发货接口 APIResponse
taobao.wlb.order.jz.consign

家装类订单使用该接口发货
*/
type TaobaoWlbOrderJzConsignAPIResponse struct {
    model.CommonResponse
    TaobaoWlbOrderJzConsignResponse
}

type TaobaoWlbOrderJzConsignResponse struct {
    XMLName xml.Name `xml:"wlb_order_jz_consign_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ResultErrorCode   string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`

    
    // 错误信息描述
    
    ResultErrorMsg   string `json:"result_error_msg,omitempty" xml:"result_error_msg,omitempty"`

    
    // 是否成功
    
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`

    
}
