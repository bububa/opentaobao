package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
客服外包订单分配的商家子账号列表 APIResponse
taobao.weike.eservice.subusers.get

获取客服外包订单分配的商家子账号列表，以及授权状态
*/
type TaobaoWeikeEserviceSubusersGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWeikeEserviceSubusersGetResponse `json:"taobao_weike_eservice_subusers_get_response,omitempty"`
}

type TaobaoWeikeEserviceSubusersGetResponse struct {

    // 商家子账号查询结果
    Result   *AuthorizedAccountWrapper `json:"result,omitempty"`

}
