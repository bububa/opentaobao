package tmallcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
isv查询服务工单详情 APIResponse
tmall.aliauto.service.receipt.get

isv查询服务工单详情
*/
type TmallAliautoServiceReceiptGetAPIResponse struct {
    model.CommonResponse
    TmallAliautoServiceReceiptGetResponse
}

type TmallAliautoServiceReceiptGetResponse struct {
    XMLName xml.Name `xml:"tmall_aliauto_service_receipt_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回包装类
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
