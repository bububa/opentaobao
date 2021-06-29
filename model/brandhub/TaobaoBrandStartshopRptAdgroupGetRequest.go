package brandhub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
明星店铺推广单元报表数据查询 APIRequest
taobao.brand.startshop.rpt.adgroup.get

获取明星店铺广告adgroup分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandStartshopRptAdgroupGetRequest struct {
    model.Params

    // 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
    trafficType   string 

    // 转化周期默认15天,3,7,15
    effect   int64 

    // 当前页数
    pageIndex   string 

    // 每页条数
    pageSize   string 

    // 开始时间(最多可查询最近90天)
    startDate   string 

    // 截至时间(最晚到昨天)
    endDate   string 

}

func NewTaobaoBrandStartshopRptAdgroupGetRequest() *TaobaoBrandStartshopRptAdgroupGetRequest{
    return &TaobaoBrandStartshopRptAdgroupGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetApiMethodName() string {
    return "taobao.brand.startshop.rpt.adgroup.get"
}

func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBrandStartshopRptAdgroupGetRequest) SetTrafficType(trafficType string) error {
    r.trafficType = trafficType
    r.Set("traffic_type", trafficType)
    return nil
}

func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetTrafficType() string {
    return r.trafficType
}

func (r *TaobaoBrandStartshopRptAdgroupGetRequest) SetEffect(effect int64) error {
    r.effect = effect
    r.Set("effect", effect)
    return nil
}

func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetEffect() int64 {
    return r.effect
}

func (r *TaobaoBrandStartshopRptAdgroupGetRequest) SetPageIndex(pageIndex string) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetPageIndex() string {
    return r.pageIndex
}

func (r *TaobaoBrandStartshopRptAdgroupGetRequest) SetPageSize(pageSize string) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetPageSize() string {
    return r.pageSize
}

func (r *TaobaoBrandStartshopRptAdgroupGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoBrandStartshopRptAdgroupGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetEndDate() string {
    return r.endDate
}

