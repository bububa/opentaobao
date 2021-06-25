package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广组实时报表数据 APIRequest
taobao.simba.rtrpt.adgroup.get

获取推广组实时报表数据
*/
type TaobaoSimbaRtrptAdgroupGetRequest struct {
    model.Params

    // 用户名
    nick   string 

    // 推广计划id
    campaignId   int64 

    // 日期，格式yyyy-mm-dd
    theDate   string 

    // 每页大小
    pageSize   int64 

    // 页码
    pageNumber   int64 

}

func NewTaobaoSimbaRtrptAdgroupGetRequest() *TaobaoSimbaRtrptAdgroupGetRequest{
    return &TaobaoSimbaRtrptAdgroupGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaRtrptAdgroupGetRequest) GetApiMethodName() string {
    return "taobao.simba.rtrpt.adgroup.get"
}

func (r TaobaoSimbaRtrptAdgroupGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaRtrptAdgroupGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaRtrptAdgroupGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaRtrptAdgroupGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaRtrptAdgroupGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaRtrptAdgroupGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

func (r TaobaoSimbaRtrptAdgroupGetRequest) GetTheDate() string {
    return r.theDate
}

func (r *TaobaoSimbaRtrptAdgroupGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaRtrptAdgroupGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaRtrptAdgroupGetRequest) SetPageNumber(pageNumber int64) error {
    r.pageNumber = pageNumber
    r.Set("page_number", pageNumber)
    return nil
}

func (r TaobaoSimbaRtrptAdgroupGetRequest) GetPageNumber() int64 {
    return r.pageNumber
}

