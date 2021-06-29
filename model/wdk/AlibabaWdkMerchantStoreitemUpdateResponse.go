package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店商品 API返回值 
alibaba.wdk.merchant.storeitem.update

修改门店商品
*/
type AlibabaWdkMerchantStoreitemUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMerchantStoreitemUpdateResponse
}

// 修改门店商品 成功返回结果
type AlibabaWdkMerchantStoreitemUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_merchant_storeitem_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaWdkMerchantStoreitemUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
