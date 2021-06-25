package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家产品服务-编辑产品 APIResponse
alibaba.wdk.merchantproduct.edit

商家产品服务-编辑产品
*/
type AlibabaWdkMerchantproductEditAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMerchantproductEditResponse `json:"alibaba_wdk_merchantproduct_edit_response,omitempty"`
}

type AlibabaWdkMerchantproductEditResponse struct {

    // 产品编辑返回结果
    Result   *AlibabaWdkMerchantproductEditApiResult `json:"result,omitempty"`

}
