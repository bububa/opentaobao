package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量得到推广计划下的推广单元 APIRequest
taobao.simba.adgroupsbycampaignid.get

根据推广计划ID分页获取推广计划下的推广单元信息
*/
type TaobaoSimbaAdgroupsbycampaignidGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 推广计划Id
    campaignId   int64 

    // 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
    pageSize   int64 

    // 页码，从1开始
    pageNo   int64 

}

func NewTaobaoSimbaAdgroupsbycampaignidGetRequest() *TaobaoSimbaAdgroupsbycampaignidGetRequest{
    return &TaobaoSimbaAdgroupsbycampaignidGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaAdgroupsbycampaignidGetRequest) GetApiMethodName() string {
    return "taobao.simba.adgroupsbycampaignid.get"
}

func (r TaobaoSimbaAdgroupsbycampaignidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaAdgroupsbycampaignidGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaAdgroupsbycampaignidGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaAdgroupsbycampaignidGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaAdgroupsbycampaignidGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaAdgroupsbycampaignidGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaAdgroupsbycampaignidGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaAdgroupsbycampaignidGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaAdgroupsbycampaignidGetRequest) GetPageNo() int64 {
    return r.pageNo
}

