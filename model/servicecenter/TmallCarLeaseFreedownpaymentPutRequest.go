package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步直租车免首付商品活动信息 API请求
tmall.car.lease.freedownpayment.put

汽车行业直租车免首付需求中，用与对商品打标，活动范围设置，在消费者端商品详情页、订单等环节透出，表示该商品为直租免首付商品。
*/
type TmallCarLeaseFreedownpaymentPutAPIRequest struct {
    model.Params
    // 活动预热结束时间:格式：yyyy.MM.dd HH:mm:ss
    _preEndTime   string
    // 活动预热开始时间:格式：yyyy.MM.dd HH:mm:ss
    _preStartTime   string
    // 商品ID
    _itemId   int64
    // 活动时间范围节点(json格式字符串)：<br/> 开始时间(startTime),格式：yyyy.MM.dd HH:mm:ss <br/>  结束时间(endTime),格式：yyyy.MM.dd HH:mm:ss <br/>  名额(amount)
    _timeRangeList   string
    // 外部活动ID
    _refActivityId   string
}

// 初始化TmallCarLeaseFreedownpaymentPutAPIRequest对象
func NewTmallCarLeaseFreedownpaymentPutRequest() *TmallCarLeaseFreedownpaymentPutAPIRequest{
    return &TmallCarLeaseFreedownpaymentPutAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetApiMethodName() string {
    return "tmall.car.lease.freedownpayment.put"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PreEndTime Setter
// 活动预热结束时间:格式：yyyy.MM.dd HH:mm:ss
func (r *TmallCarLeaseFreedownpaymentPutAPIRequest) SetPreEndTime(_preEndTime string) error {
    r._preEndTime = _preEndTime
    r.Set("pre_end_time", _preEndTime)
    return nil
}

// PreEndTime Getter
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetPreEndTime() string {
    return r._preEndTime
}
// PreStartTime Setter
// 活动预热开始时间:格式：yyyy.MM.dd HH:mm:ss
func (r *TmallCarLeaseFreedownpaymentPutAPIRequest) SetPreStartTime(_preStartTime string) error {
    r._preStartTime = _preStartTime
    r.Set("pre_start_time", _preStartTime)
    return nil
}

// PreStartTime Getter
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetPreStartTime() string {
    return r._preStartTime
}
// ItemId Setter
// 商品ID
func (r *TmallCarLeaseFreedownpaymentPutAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetItemId() int64 {
    return r._itemId
}
// TimeRangeList Setter
// 活动时间范围节点(json格式字符串)：<br/> 开始时间(startTime),格式：yyyy.MM.dd HH:mm:ss <br/>  结束时间(endTime),格式：yyyy.MM.dd HH:mm:ss <br/>  名额(amount)
func (r *TmallCarLeaseFreedownpaymentPutAPIRequest) SetTimeRangeList(_timeRangeList string) error {
    r._timeRangeList = _timeRangeList
    r.Set("time_range_list", _timeRangeList)
    return nil
}

// TimeRangeList Getter
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetTimeRangeList() string {
    return r._timeRangeList
}
// RefActivityId Setter
// 外部活动ID
func (r *TmallCarLeaseFreedownpaymentPutAPIRequest) SetRefActivityId(_refActivityId string) error {
    r._refActivityId = _refActivityId
    r.Set("ref_activity_id", _refActivityId)
    return nil
}

// RefActivityId Getter
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetRefActivityId() string {
    return r._refActivityId
}
