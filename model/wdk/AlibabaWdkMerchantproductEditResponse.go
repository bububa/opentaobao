package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家产品服务-编辑产品 APIResponse
alibaba.wdk.merchantproduct.edit

商家产品服务-编辑产品
*/
type AlibabaWdkMerchantproductEditAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_merchantproduct_edit_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 产品编辑返回结果
    
    Result   *AlibabaWdkMerchantproductEditApiResult `json:"result,omitempty" xml:"