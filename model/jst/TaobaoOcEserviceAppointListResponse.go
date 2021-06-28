package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
交互卡片预约信息读取接口 APIResponse
taobao.oc.eservice.appoint.list

允许外部的isv通过该接口读取门店预约信息
*/
type TaobaoOcEserviceAppointListAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOcEserviceAppointListResponse `json:"oc_eservice_appoint_list_response,omitempty"` 
    TaobaoOcEserviceAppointListResponse
}

/* model for simplify = false
type TaobaoOcEserviceAppointListResponse struct {

    // 返回的预约信息，多个预约信息按照预约时间升序排序
    
    Models  struct {
        O2oAppointInfoDto  []O2oAppointInfoDto `json:"o2o_appoint_info_dto,omitempty"`
    } `json:"models,omitempty"`
    

    // 返回的预约信息数目
    
    TotalNumber   int64 `json:"total_number,omitempty"`
    

}
*/

type TaobaoOcEserviceAppointListResponse struct {

    // 返回的预约信息，多个预约信息按照预约时间升序排序
    Models   []O2oAppointInfoDto `json:"models,omitempty"`

    // 返回的预约信息数目
    TotalNumber   int64 `json:"total_number,omitempty"`

}
