package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家类目修改接口 APIResponse
alibaba.wdk.sku.category.update

商家类目修改接口
*/
type AlibabaWdkSkuCategoryUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuCategoryUpdateResponse
}

type AlibabaWdkSkuCategoryUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_category_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用结果
    
    Result   *AlibabaWdkSkuCategoryUpdateApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
