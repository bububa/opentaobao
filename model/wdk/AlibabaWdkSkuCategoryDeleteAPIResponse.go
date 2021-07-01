package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家类目删除接口 API返回值 
alibaba.wdk.sku.category.delete

商家类目删除接口
*/
type AlibabaWdkSkuCategoryDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuCategoryDeleteAPIResponseModel
}

// 商家类目删除接口 成功返回结果
type AlibabaWdkSkuCategoryDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_category_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用结果
    Result   *AlibabaWdkSkuCategoryDeleteApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
