package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询当天可添加的余量 APIResponse
taobao.baichuan.item.subscribe.daily.left.query

查询当天可添加的余量
*/
type TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBaichuanItemSubscribeDailyLeftQueryResponse `json:"baichuan_item_subscribe_daily_left_query_response,omitempty"` 
    TaobaoBaichuanItemSubscribeDailyLeftQueryResponse
}

/* model for simplify = false
type TaobaoBaichuanItemSubscribeDailyLeftQueryResponse struct {

    // 接口返回model
    
    Result  *struct {
        TaobaoBaichuanItemSubscribeDailyLeftQueryResult  *TaobaoBaichuanItemSubscribeDailyLeftQueryResult `json:"taobao_baichuan_item_subscribe_daily_left_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoBaichuanItemSubscribeDailyLeftQueryResponse struct {

    // 接口返回model
    Result   *TaobaoBaichuanItemSubscribeDailyLeftQueryResult `json:"result,omitempty"`

}
