package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步直租车免首付商品活动信息 APIRequest
tmall.car.lease.freedownpayment.put

汽车行业直租车免首付需求中，用与对商品打标，活动范围设置，在消费者端商品详情页、订单等环节透出，表示该商品为直租免首付商品。
*/
type TmallCarLeaseFreedownpaymentPutRequest struct {
    model.Params

    // 活动预热结束时间:格式：yyyy.MM.dd HH:mm:ss
    preEndTime   string 

    // 活动预热开始时间:格式：yyyy.MM.dd HH:mm:ss
    preStartTime   string 

    // 商品ID
    itemId   int64 

    // 活动时间范围节点(json格式字符串)：<br/> 开始时间(startTime),格式：yyyy.MM.dd HH:mm:ss <br/>  结束时间(endTime),格式：yyyy.MM.dd HH:mm:ss <br/>  名额(amount)
    timeRangeList   string 

    // 外部活动ID
    refActivityId   string 

}

func NewTmallCarLeaseFreedownpaymentPutRequest() *TmallCarLeaseFreedownpaymentPutRequest{
    return &TmallCarLeaseFreedownpaymentPutRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarLeaseFreedownpaymentPutRequest) GetApiMethodName() string {
    return "tmall.car.lease.freedownpayment.put"
}

func (r TmallCarLeaseFreedownpaymentPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarLeaseFreedownpaymentPutRequest) SetPreEndTime(preEndTime string) error {
    r.preEndTime = preEndTime
    r.Set("pre_end_time", preEndTime)
    return nil
}

func (r TmallCarLeaseFreedownpaymentPutRequest) GetPreEndTime() string {
    return r.preEndTime
}

func (r *TmallCarLeaseFreedownpaymentPutRequest) SetPreStartTime(preStartTime string) error {
    r.preStartTime = preStartTime
    r.Set("pre_start_time", preStartTime)
    return nil
}

func (r TmallCarLeaseFreedownpaymentPutRequest) GetPreStartTime() string {
    return r.preStartTime
}

func (r *TmallCarLeaseFreedownpaymentPutRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallCarLeaseFreedownpaymentPutRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TmallCarLeaseFreedownpaymentPutRequest) SetTimeRangeList(timeRangeList string) error {
    r.timeRangeList = timeRangeList
    r.Set("time_range_list", timeRangeList)
    return nil
}

func (r TmallCarLeaseFreedownpaymentPutRequest) GetTimeRangeList() string {
    return r.timeRangeList
}

func (r *TmallCarLeaseFreedownpaymentPutRequest) SetRefActivityId(refActivityId string) error {
    r.refActivityId = refActivityId
    r.Set("ref_activity_id", refActivityId)
    return nil
}

func (r TmallCarLeaseFreedownpaymentPutRequest) GetRefActivityId() string {
    return r.refActivityId
}

