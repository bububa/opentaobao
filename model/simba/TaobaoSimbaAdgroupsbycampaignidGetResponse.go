package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量得到推广计划下的推广单元 API返回值 
taobao.simba.adgroupsbycampaignid.get

根据推广计划ID分页获取推广计划下的推广单元信息
*/
type TaobaoSimbaAdgroupsbycampaignidGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaAdgroupsbycampaignidGetResponse
}

// 批量得到推广计划下的推广单元 成功返回结果
type TaobaoSimbaAdgroupsbycampaignidGetResponse struct {
    XMLName xml.Name `xml:"simba_adgroupsbycampaignid_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回的推广组分页对象
    Adgroups   *ADGroupPage `json:"adgroups,omitempty" xml:"adgroups,omitempty"`
}
