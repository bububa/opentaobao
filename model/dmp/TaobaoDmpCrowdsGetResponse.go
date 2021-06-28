package dmp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询人群服务 APIResponse
taobao.dmp.crowds.get

查询人群服务
*/
type TaobaoDmpCrowdsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoDmpCrowdsGetResponse `json:"dmp_crowds_get_response,omitempty"` 
    TaobaoDmpCrowdsGetResponse
}

/* model for simplify = false
type TaobaoDmpCrowdsGetResponse struct {

    // 1
    
    Results  struct {
        DmpCrowdDTO  []DmpCrowdDTO `json:"dmp_crowd_dto,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type TaobaoDmpCrowdsGetResponse struct {

    // 1
    Results   []DmpCrowdDTO `json:"results,omitempty"`

}
