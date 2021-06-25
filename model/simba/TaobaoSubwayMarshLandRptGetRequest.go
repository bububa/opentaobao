package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取捡漏词包分时报表数据 APIRequest
taobao.subway.marsh.land.rpt.get

获取捡漏词包分时报表数据
*/
type TaobaoSubwayMarshLandRptGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 2021-05-11
    endDate   string 

    // 推广组id
    adgroupIdEqual   string 

    // 词包类型（捡漏词包填19）
    isAutoMatchEqual   string 

    // 计划id
    campaignIdEqual   string 

    // 2021-05-05
    startDate   string 

}

func NewTaobaoSubwayMarshLandRptGetRequest() *TaobaoSubwayMarshLandRptGetRequest{
    return &TaobaoSubwayMarshLandRptGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSubwayMarshLandRptGetRequest) GetApiMethodName() string {
    return "taobao.subway.marsh.land.rpt.get"
}

func (r TaobaoSubwayMarshLandRptGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSubwayMarshLandRptGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSubwayMarshLandRptGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSubwayMarshLandRptGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoSubwayMarshLandRptGetRequest) GetEndDate() string {
    return r.endDate
}

func (r *TaobaoSubwayMarshLandRptGetRequest) SetAdgroupIdEqual(adgroupIdEqual string) error {
    r.adgroupIdEqual = adgroupIdEqual
    r.Set("adgroup_id_equal", adgroupIdEqual)
    return nil
}

func (r TaobaoSubwayMarshLandRptGetRequest) GetAdgroupIdEqual() string {
    return r.adgroupIdEqual
}

func (r *TaobaoSubwayMarshLandRptGetRequest) SetIsAutoMatchEqual(isAutoMatchEqual string) error {
    r.isAutoMatchEqual = isAutoMatchEqual
    r.Set("is_auto_match_equal", isAutoMatchEqual)
    return nil
}

func (r TaobaoSubwayMarshLandRptGetRequest) GetIsAutoMatchEqual() string {
    return r.isAutoMatchEqual
}

func (r *TaobaoSubwayMarshLandRptGetRequest) SetCampaignIdEqual(campaignIdEqual string) error {
    r.campaignIdEqual = campaignIdEqual
    r.Set("campaign_id_equal", campaignIdEqual)
    return nil
}

func (r TaobaoSubwayMarshLandRptGetRequest) GetCampaignIdEqual() string {
    return r.campaignIdEqual
}

func (r *TaobaoSubwayMarshLandRptGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoSubwayMarshLandRptGetRequest) GetStartDate() string {
    return r.startDate
}

