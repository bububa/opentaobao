package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseFreedownpaymentPutAPIRequest 同步直租车免首付商品活动信息 API请求
// tmall.car.lease.freedownpayment.put
//
// 汽车行业直租车免首付需求中，用与对商品打标，活动范围设置，在消费者端商品详情页、订单等环节透出，表示该商品为直租免首付商品。
type TmallCarLeaseFreedownpaymentPutAPIRequest struct {
	model.Params
	// 活动预热结束时间:格式：yyyy.MM.dd HH:mm:ss
	_preEndTime string
	// 活动预热开始时间:格式：yyyy.MM.dd HH:mm:ss
	_preStartTime string
	// 活动时间范围节点(json格式字符串)：<br/> 开始时间(startTime),格式：yyyy.MM.dd HH:mm:ss <br/>  结束时间(endTime),格式：yyyy.MM.dd HH:mm:ss <br/>  名额(amount)
	_timeRangeList string
	// 外部活动ID
	_refActivityId string
	// 商品ID
	_itemId int64
}

// NewTmallCarLeaseFreedownpaymentPutRequest 初始化TmallCarLeaseFreedownpaymentPutAPIRequest对象
func NewTmallCarLeaseFreedownpaymentPutRequest() *TmallCarLeaseFreedownpaymentPutAPIRequest {
	return &TmallCarLeaseFreedownpaymentPutAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarLeaseFreedownpaymentPutAPIRequest) Reset() {
	r._preEndTime = ""
	r._preStartTime = ""
	r._timeRangeList = ""
	r._refActivityId = ""
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.freedownpayment.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreEndTime is PreEndTime Setter
// 活动预热结束时间:格式：yyyy.MM.dd HH:mm:ss
func (r *TmallCarLeaseFreedownpaymentPutAPIRequest) SetPreEndTime(_preEndTime string) error {
	r._preEndTime = _preEndTime
	r.Set("pre_end_time", _preEndTime)
	return nil
}

// GetPreEndTime PreEndTime Getter
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetPreEndTime() string {
	return r._preEndTime
}

// SetPreStartTime is PreStartTime Setter
// 活动预热开始时间:格式：yyyy.MM.dd HH:mm:ss
func (r *TmallCarLeaseFreedownpaymentPutAPIRequest) SetPreStartTime(_preStartTime string) error {
	r._preStartTime = _preStartTime
	r.Set("pre_start_time", _preStartTime)
	return nil
}

// GetPreStartTime PreStartTime Getter
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetPreStartTime() string {
	return r._preStartTime
}

// SetTimeRangeList is TimeRangeList Setter
// 活动时间范围节点(json格式字符串)：&lt;br/&gt; 开始时间(startTime),格式：yyyy.MM.dd HH:mm:ss &lt;br/&gt;  结束时间(endTime),格式：yyyy.MM.dd HH:mm:ss &lt;br/&gt;  名额(amount)
func (r *TmallCarLeaseFreedownpaymentPutAPIRequest) SetTimeRangeList(_timeRangeList string) error {
	r._timeRangeList = _timeRangeList
	r.Set("time_range_list", _timeRangeList)
	return nil
}

// GetTimeRangeList TimeRangeList Getter
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetTimeRangeList() string {
	return r._timeRangeList
}

// SetRefActivityId is RefActivityId Setter
// 外部活动ID
func (r *TmallCarLeaseFreedownpaymentPutAPIRequest) SetRefActivityId(_refActivityId string) error {
	r._refActivityId = _refActivityId
	r.Set("ref_activity_id", _refActivityId)
	return nil
}

// GetRefActivityId RefActivityId Getter
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetRefActivityId() string {
	return r._refActivityId
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallCarLeaseFreedownpaymentPutAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallCarLeaseFreedownpaymentPutAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTmallCarLeaseFreedownpaymentPutAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarLeaseFreedownpaymentPutRequest()
	},
}

// GetTmallCarLeaseFreedownpaymentPutRequest 从 sync.Pool 获取 TmallCarLeaseFreedownpaymentPutAPIRequest
func GetTmallCarLeaseFreedownpaymentPutAPIRequest() *TmallCarLeaseFreedownpaymentPutAPIRequest {
	return poolTmallCarLeaseFreedownpaymentPutAPIRequest.Get().(*TmallCarLeaseFreedownpaymentPutAPIRequest)
}

// ReleaseTmallCarLeaseFreedownpaymentPutAPIRequest 将 TmallCarLeaseFreedownpaymentPutAPIRequest 放入 sync.Pool
func ReleaseTmallCarLeaseFreedownpaymentPutAPIRequest(v *TmallCarLeaseFreedownpaymentPutAPIRequest) {
	v.Reset()
	poolTmallCarLeaseFreedownpaymentPutAPIRequest.Put(v)
}
