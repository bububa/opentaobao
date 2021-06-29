package brandhub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
明星店铺品牌流量包报表数据查询 APIRequest
taobao.brand.startshop.rpt.wordpackage.get

获取明星店铺广告词包分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandStartshopRptWordpackageGetRequest struct {
    model.Params

    // 开始日期
    startDate   string 

    // 结束日期
    endDate   string 

    // 转化周期
    effect   string 

    // 流量类型
    trafficType   string 

    // 每页显示条数(0,200]
    pageSize   string 

    // 当前页数 ,从1开始
    pageIndex   string 

}

func NewTaobaoBrandStartshopRptWordpackageGetRequest() *TaobaoBrandStartshopRptWordpackageGetRequest{
    return &TaobaoBrandStartshopRptWordpackageGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBrandStartshopRptWordpackageGetRequest) GetApiMethodName() string {
    return "taobao.brand.startshop.rpt.wordpackage.get"
}

func (r TaobaoBrandStartshopRptWordpackageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBrandStartshopRptWordpackageGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoBrandStartshopRptWordpackageGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoBrandStartshopRptWordpackageGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoBrandStartshopRptWordpackageGetRequest) GetEndDate() string {
    return r.endDate
}

func (r *TaobaoBrandStartshopRptWordpackageGetRequest) SetEffect(effect string) error {
    r.effect = effect
    r.Set("effect", effect)
    return nil
}

func (r TaobaoBrandStartshopRptWordpackageGetRequest) GetEffect() string {
    return r.effect
}

func (r *TaobaoBrandStartshopRptWordpackageGetRequest) SetTrafficType(trafficType string) error {
    r.trafficType = trafficType
    r.Set("traffic_type", trafficType)
    return nil
}

func (r TaobaoBrandStartshopRptWordpackageGetRequest) GetTrafficType() string {
    return r.trafficType
}

func (r *TaobaoBrandStartshopRptWordpackageGetRequest) SetPageSize(pageSize string) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoBrandStartshopRptWordpackageGetRequest) GetPageSize() string {
    return r.pageSize
}

func (r *TaobaoBrandStartshopRptWordpackageGetRequest) SetPageIndex(pageIndex string) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r TaobaoBrandStartshopRptWordpackageGetRequest) GetPageIndex() string {
    return r.pageIndex
}

