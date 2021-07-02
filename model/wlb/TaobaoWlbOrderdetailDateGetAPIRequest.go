package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderdetailDateGetAPIRequest 按照日期范围查询物流订单详情 API请求
// taobao.wlb.orderdetail.date.get
//
// 外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情
type TaobaoWlbOrderdetailDateGetAPIRequest struct {
	model.Params
	// 创建时间起始
	_startTime string
	// 创建时间结束
	_endTime string
	// 分页大小
	_pageSize int64
	// 分页下标
	_pageNo int64
}

// NewTaobaoWlbOrderdetailDateGetRequest 初始化TaobaoWlbOrderdetailDateGetAPIRequest对象
func NewTaobaoWlbOrderdetailDateGetRequest() *TaobaoWlbOrderdetailDateGetAPIRequest {
	return &TaobaoWlbOrderdetailDateGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderdetailDateGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.orderdetail.date.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderdetailDateGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStartTime is StartTime Setter
// 创建时间起始
func (r *TaobaoWlbOrderdetailDateGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoWlbOrderdetailDateGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 创建时间结束
func (r *TaobaoWlbOrderdetailDateGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoWlbOrderdetailDateGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TaobaoWlbOrderdetailDateGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoWlbOrderdetailDateGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 分页下标
func (r *TaobaoWlbOrderdetailDateGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoWlbOrderdetailDateGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
