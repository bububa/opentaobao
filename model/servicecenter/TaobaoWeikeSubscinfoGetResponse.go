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
    Response *TaobaoWeikeSubscinfoGetResponse `json:"taobao_weike_subscinfo_get_response,omitempty"`
}

type TaobaoWeikeSubscinfoGetResponse struct {

    // 返回结果
    Result   *SubscInfoWrapper `json:"result,omitempty"`

}
