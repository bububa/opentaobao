package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存订单投保后回传保单号 APIResponse
alibaba.wdkorder.sharestock.insurance.callback

共享库存订单投保消息获取
*/
type AlibabaWdkorderSharestockInsuranceCallbackAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdkorder_sharestock_insurance_callback_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    Result   *MaochaoOrderInsuranceCallbackResult `json:"result,omitempty" xml:"