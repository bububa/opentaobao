package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
系外拉新代理商增加客户经理接口 APIResponse
alibaba.aliqin.offline.customer.add

阿里通信这边维护了代理商和其对应的客户经理的关系，用于业务处理，开放该接口用于代理商将他们系统下的客户经理信息同步给我们
*/
type AlibabaAliqinOfflineCustomerAddAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinOfflineCustomerAddResponse `json:"alibaba_aliqin_offline_customer_add_response,omitempty"` 
    AlibabaAliqinOfflineCustomerAddResponse
}

/* model for simplify = false
type AlibabaAliqinOfflineCustomerAddResponse struct {

    // result
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinOfflineCustomerAddResponse struct {

    // result
    Result   *CommonResult `json:"result,omitempty"`

}
