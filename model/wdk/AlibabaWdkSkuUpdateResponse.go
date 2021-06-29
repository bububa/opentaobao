package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品 APIResponse
alibaba.wdk.sku.update

开放商品更新接口
*/
type AlibabaWdkSkuUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuUpdateResponse
}

type AlibabaWdkSkuUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 执行结果
    
    Result   *AlibabaWdkSkuUpdateApiResults `json:"result,omitempty" xml:"result,omitempty"`

    
}
