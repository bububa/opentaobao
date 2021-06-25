package jms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询ONS分组 APIResponse
taobao.jushita.jms.group.get

查询当前appkey在ONS中已有的分组
*/
type TaobaoJushitaJmsGroupGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJushitaJmsGroupGetResponse `json:"taobao_jushita_jms_group_get_response,omitempty"`
}

type TaobaoJushitaJmsGroupGetResponse struct {

    // 返回的总数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 分组信息
    Groups   []MsgGroupDO `json:"groups,omitempty"`

}
