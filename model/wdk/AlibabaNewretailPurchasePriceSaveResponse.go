package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存 采购价上传接口 APIResponse
alibaba.newretail.purchase.price.save

共享库存业务 供应商上传商品采购价
*/
type AlibabaNewretailPurchasePriceSaveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_newretail_purchase_price_save_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果对象
    
    Result   *TopBaseResult `json:"result,omitempty" xml:"