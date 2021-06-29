package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品删除 APIResponse
alibaba.item.operate.delete

商品删除
*/
type AlibabaItemOperateDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaItemOperateDeleteResponse
}

type AlibabaItemOperateDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_item_operate_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 商品删除是否成功
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
