package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据屏编号获取专柜集 APIResponse
alibaba.mos.store.getstorelist

根据屏编号获取专柜集
*/
type AlibabaMosStoreGetstorelistAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMosStoreGetstorelistResponse `json:"alibaba_mos_store_getstorelist_response,omitempty"` 
    AlibabaMosStoreGetstorelistResponse
}

/* model for simplify = false
type AlibabaMosStoreGetstorelistResponse struct {

    // result
    
    Result  *struct {
        AlibabaMosStoreGetstorelistResultDo  *AlibabaMosStoreGetstorelistResultDo `json:"alibaba_mos_store_getstorelist_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaMosStoreGetstorelistResponse struct {

    // result
    Result   *AlibabaMosStoreGetstorelistResultDo `json:"result,omitempty"`

}
