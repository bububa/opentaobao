package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增商品 APIResponse
alibaba.wdk.sku.add

创建RT门店商品或DC商品
*/
type AlibabaWdkSkuAddAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuAddResponse
}

type AlibabaWdkSkuAddResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    Result   *AlibabaWdkSkuAddApiResults `json:"result,omitempty" xml:"result,omitempty"`

    
}
