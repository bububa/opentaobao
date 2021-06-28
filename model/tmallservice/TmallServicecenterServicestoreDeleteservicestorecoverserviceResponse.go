package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除网点覆盖的服务 APIResponse
tmall.servicecenter.servicestore.deleteservicestorecoverservice

天猫服务平台删除网点覆盖的服务，
必选字段：serviceStoreCode、bizType
*/
type TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterServicestoreDeleteservicestorecoverserviceResponse `json:"tmall_servicecenter_servicestore_deleteservicestorecoverservice_response,omitempty"` 
    TmallServicecenterServicestoreDeleteservicestorecoverserviceResponse
}

/* model for simplify = false
type TmallServicecenterServicestoreDeleteservicestorecoverserviceResponse struct {

    // result
    
    Result  *struct {
        ResultBase  *ResultBase `json:"result_base,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterServicestoreDeleteservicestorecoverserviceResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
