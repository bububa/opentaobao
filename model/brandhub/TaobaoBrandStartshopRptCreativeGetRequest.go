package brandhub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
明星店铺创意报表数据查询 APIRequest
taobao.brand.startshop.rpt.creative.get

获取明星店铺广告creative分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandStartshopRptCreativeGetRequest struct {
    model.Params

    // 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
    trafficType   string 

    // 当前页数
    pageIndex   string 

    // 每页条数
    pageSize   string 

    // 转化周期,默认15, 3,7,15
    effect   string 

    // 开始日期(最多查询1个月的数据)
    startDate   string 

    // 截至日期(最晚到昨天)
    endDate   string 

}

func NewTaobaoBrandStartshopRptCreativeGetRequest() *TaobaoBrandStartshopRptCreativeGetRequest{
    return &TaobaoBrandStartshopRptCreativeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBrandStartshopRptCreativeGetRequest) GetApiMethodName() string {
    return "taobao.brand.startshop.rpt.creative.get"
}

func (r TaobaoBrandStartshopRptCreativeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBrandStartshopRptCreativeGetRequest) SetTrafficType(trafficType string) error {
    r.trafficType = trafficType
    r.Set("traffic_type", trafficType)
    return nil
}

func (r TaobaoBrandStartshopRptCreativeGetRequest) GetTrafficType() string {
    return r.trafficType
}

func (r *TaobaoBrandStartshopRptCreativeGetRequest) SetPageIndex(pageIndex string) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r TaobaoBrandStartshopRptCreativeGetRequest) GetPageIndex() string {
    return r.pageIndex
}

func (r *TaobaoBrandStartshopRptCreativeGetRequest) SetPageSize(pageSize string) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoBrandStartshopRptCreativeGetRequest) GetPageSize() string {
    return r.pageSize
}

func (r *TaobaoBrandStartshopRptCreativeGetRequest) SetEffect(effect string) error {
    r.effect = effect
    r.Set("effect", effect)
    return nil
}

func (r TaobaoBrandStartshopRptCreativeGetRequest) GetEffect() string {
    return r.effect
}

func (r *TaobaoBrandStartshopRptCreativeGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoBrandStartshopRptCreativeGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoBrandStartshopRptCreativeGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoBrandStartshopRptCreativeGetRequest) GetEndDate() string {
    return r.endDate
}

