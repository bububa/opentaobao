package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
预售商家仓接单 APIResponse
alibaba.ascp.uop.taobao.presalesorder.create

预售商家仓接单
*/
type AlibabaAscpUopTaobaoPresalesorderCreateAPIResponse struct {
    model.CommonResponse
    AlibabaAscpUopTaobaoPresalesorderCreateResponse
}

type AlibabaAscpUopTaobaoPresalesorderCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_uop_taobao_presalesorder_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值包装,result为返回具体消息内容
    
    PresalesOrderCreateResponse   *ResultWrapper `json:"presales_order_create_response,omitempty" xml:"presales_order_create_response,omitempty"`

    
}
