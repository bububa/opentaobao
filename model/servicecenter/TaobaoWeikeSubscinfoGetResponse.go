package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
需求订单查询接口 APIResponse
taobao.weike.subscinfo.get

需求订单查询接口
*/
type TaobaoWeikeSubscinfoGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWeikeSubscinfoGetResponse `json:"weike_subscinfo_get_response,omitempty"` 
    TaobaoWeikeSubscinfoGetResponse
}

/* model for simplify = false
type TaobaoWeikeSubscinfoGetResponse struct {

    // 返回结果
    
    Result  *struct {
        SubscInfoWrapper  *SubscInfoWrapper `json:"subsc_info_wrapper,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWeikeSubscinfoGetResponse struct {

    // 返回结果
    Result   *SubscInfoWrapper `json:"result,omitempty"`

}
