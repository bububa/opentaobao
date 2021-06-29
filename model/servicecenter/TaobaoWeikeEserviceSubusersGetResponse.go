package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
客服外包订单分配的商家子账号列表 API返回值 
taobao.weike.eservice.subusers.get

获取客服外包订单分配的商家子账号列表，以及授权状态
*/
type TaobaoWeikeEserviceSubusersGetAPIResponse struct {
    model.CommonResponse
    TaobaoWeikeEserviceSubusersGetResponse
}

// 客服外包订单分配的商家子账号列表 成功返回结果
type TaobaoWeikeEserviceSubusersGetResponse struct {
    XMLName xml.Name `xml:"weike_eservice_subusers_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 商家子账号查询结果
    Result   *AuthorizedAccountWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
