package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取开放平台出口IP段 APIResponse
taobao.top.ipout.get

获取开放平台出口IP段
*/
type TaobaoTopIpoutGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTopIpoutGetResponse `json:"taobao_top_ipout_get_response,omitempty"`
}

type TaobaoTopIpoutGetResponse struct {

    // TOP网关出口IP列表
    IpList   string `json:"ip_list,omitempty"`

}
