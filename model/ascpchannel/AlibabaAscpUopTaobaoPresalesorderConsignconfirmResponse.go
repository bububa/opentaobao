package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
预售商家仓出库 APIResponse
alibaba.ascp.uop.taobao.presalesorder.consignconfirm

预售商家仓出库
*/
type AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse struct {
    model.CommonResponse
    AlibabaAscpUopTaobaoPresalesorderConsignconfirmResponse
}

type AlibabaAscpUopTaobaoPresalesorderConsignconfirmResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_uop_taobao_presalesorder_consignconfirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值包装,result为返回具体消息内容
    
    PresalesOrderConsignConfirmResponse   *ResultWrapper `json:"presales_order_consign_confirm_response,omitempty" xml:"presales_order_consign_confirm_response,omitempty"`

    
}
