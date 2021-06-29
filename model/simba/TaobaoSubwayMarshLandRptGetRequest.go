package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取捡漏词包分时报表数据 API请求
taobao.subway.marsh.land.rpt.get

获取捡漏词包分时报表数据
*/
type TaobaoSubwayMarshLandRptGetRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 2021-05-11
    _endDate   string
    // 推广组id
    _adgroupIdEqual   string
    // 词包类型（捡漏词包填19）
    _isAutoMatchEqual   string
    // 计划id
    _campaignIdEqual   string
    // 2021-05-05
    _startDate   string
}

// 初始化TaobaoSubwayMarshLandRptGetRequest对象
func NewTaobaoSubwayMarshLandRptGetRequest() *TaobaoSubwayMarshLandRptGetRequest{
    return &TaobaoSubwayMarshLandRptGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSubwayMarshLandRptGetRequest) GetApiMethodName() string {
    return "taobao.subway.marsh.land.rpt.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSubwayMarshLandRptGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSubwayMarshLandRptGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSubwayMarshLandRptGetRequest) GetNick() string {
    return r._nick
}
// EndDate Setter
// 2021-05-11
func (r *TaobaoSubwayMarshLandRptGetRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoSubwayMarshLandRptGetRequest) GetEndDate() string {
    return r._endDate
}
// AdgroupIdEqual Setter
// 推广组id
func (r *TaobaoSubwayMarshLandRptGetRequest) SetAdgroupIdEqual(_adgroupIdEqual string) error {
    r._adgroupIdEqual = _adgroupIdEqual
    r.Set("adgroup_id_equal", _adgroupIdEqual)
    return nil
}

// AdgroupIdEqual Getter
func (r TaobaoSubwayMarshLandRptGetRequest) GetAdgroupIdEqual() string {
    return r._adgroupIdEqual
}
// IsAutoMatchEqual Setter
// 词包类型（捡漏词包填19）
func (r *TaobaoSubwayMarshLandRptGetRequest) SetIsAutoMatchEqual(_isAutoMatchEqual string) error {
    r._isAutoMatchEqual = _isAutoMatchEqual
    r.Set("is_auto_match_equal", _isAutoMatchEqual)
    return nil
}

// IsAutoMatchEqual Getter
func (r TaobaoSubwayMarshLandRptGetRequest) GetIsAutoMatchEqual() string {
    return r._isAutoMatchEqual
}
// CampaignIdEqual Setter
// 计划id
func (r *TaobaoSubwayMarshLandRptGetRequest) SetCampaignIdEqual(_campaignIdEqual string) error {
    r._campaignIdEqual = _campaignIdEqual
    r.Set("campaign_id_equal", _campaignIdEqual)
    return nil
}

// CampaignIdEqual Getter
func (r TaobaoSubwayMarshLandRptGetRequest) GetCampaignIdEqual() string {
    return r._campaignIdEqual
}
// StartDate Setter
// 2021-05-05
func (r *TaobaoSubwayMarshLandRptGetRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoSubwayMarshLandRptGetRequest) GetStartDate() string {
    return r._startDate
}
