package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
区域零售订单获取取件码 APIResponse
tmall.nr.order.logis.info

区域零售订单获取取件码，方便商家系统接入，获取取件码打印信息进行打印。
*/
type TmallNrOrderLogisInfoAPIResponse struct {
    model.CommonResponse
    TmallNrOrderLogisInfoResponse
}

type TmallNrOrderLogisInfoResponse struct {
    XMLName xml.Name `xml:"tmall_nr_order_logis_info_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果实体
    
    Result   *NewRetailResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
