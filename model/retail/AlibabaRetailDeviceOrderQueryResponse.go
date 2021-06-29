package retail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机订单查询 APIResponse
alibaba.retail.device.order.query

贩卖机订单查询
*/
type AlibabaRetailDeviceOrderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaRetailDeviceOrderQueryResponse
}

type AlibabaRetailDeviceOrderQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_device_order_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码
    
    Errcode   string `json:"errcode,omitempty" xml:"errcode,omitempty"`

    
    // 系统自动生成
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
    // 是否成功
    
    Data   *PaginationDo `json:"data,omitempty" xml:"data,omitempty"`

    
}
