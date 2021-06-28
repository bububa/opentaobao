package tmallsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增网点覆盖的服务 APIResponse
tmall.servicecenter.servicestore.createservicestorecoverservice

新增网点覆盖的服务，唯一性校验：服务商淘宝账号+网点编码+biz_type
前提是网点要存在，
如果需要新增的网点覆盖的服务已存在，会新增失败。
网点覆盖的服务包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、授权的类目和品牌
*/
type TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterServicestoreCreateservicestorecoverserviceResponse `json:"tmall_servicecenter_servicestore_createservicestorecoverservice_response,omitempty"` 
    TmallServicecenterServicestoreCreateservicestorecoverserviceResponse
}

/* model for simplify = false
type TmallServicecenterServicestoreCreateservicestorecoverserviceResponse struct {

    // result
    
    Result  *struct {
        ResultBase  *ResultBase `json:"result_base,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterServicestoreCreateservicestorecoverserviceResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
