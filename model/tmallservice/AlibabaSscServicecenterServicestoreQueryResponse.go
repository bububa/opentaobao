package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据天猫id查询门店信息 APIResponse
alibaba.ssc.servicecenter.servicestore.query

根据天猫id查询门店信息
*/
type AlibabaSscServicecenterServicestoreQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSscServicecenterServicestoreQueryResponse `json:"alibaba_ssc_servicecenter_servicestore_query_response,omitempty"`
}

type AlibabaSscServicecenterServicestoreQueryResponse struct {

    // 接口返回model
    Result   *AlibabaSscServicecenterServicestoreQueryResult `json:"result,omitempty"`

}
