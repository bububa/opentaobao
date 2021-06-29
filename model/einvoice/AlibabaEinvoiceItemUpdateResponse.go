package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改商品开票信息 APIResponse
alibaba.einvoice.item.update

ERP通过接口将商品的开票信息同步给阿里发票平台，自动开票时将读取这些开票信息，需要联系阿里小二开通对应的权限
*/
type AlibabaEinvoiceItemUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceItemUpdateResponse
}

type AlibabaEinvoiceItemUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_item_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 修改结果
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
