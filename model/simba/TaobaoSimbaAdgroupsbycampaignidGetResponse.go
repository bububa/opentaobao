package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量得到推广计划下的推广单元 APIResponse
taobao.simba.adgroupsbycampaignid.get

根据推广计划ID分页获取推广计划下的推广单元信息
*/
type TaobaoSimbaAdgroupsbycampaignidGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaAdgroupsbycampaignidGetResponse `json:"taobao_simba_adgroupsbycampaignid_get_response,omitempty"`
}

type TaobaoSimbaAdgroupsbycampaignidGetResponse struct {

    // 返回的推广组分页对象
    Adgroups   *ADGroupPage `json:"adgroups,omitempty"`

}
