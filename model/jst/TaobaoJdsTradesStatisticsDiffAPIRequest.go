package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojdstradesstatisticsdiffAPIRequest 订单全链路状态统计差异比较 API请求
// taobao.jds.trades.statistics.diff
//
// 订单全链路状态统计差异比较
type TaobaojdstradesstatisticsdiffAPIRequest struct {
	model.Params
	// 需要比较的状态
	_postStatus string
	// 需要比较的状态，将会和post_status做比较
	_preStatus string
	// 查询的日期，格式如YYYYMMDD的日期对应的数字
	_date int64
	// 页数
	_pageNo int64
}

// NewTaobaojdstradesstatisticsdiffRequest 初始化TaobaojdstradesstatisticsdiffAPIRequest对象
func NewTaobaojdstradesstatisticsdiffRequest() *TaobaojdstradesstatisticsdiffAPIRequest {
	return &TaobaojdstradesstatisticsdiffAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojdstradesstatisticsdiffAPIRequest) GetApiMethodName() string {
	return "taobao.jds.trades.statistics.diff"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojdstradesstatisticsdiffAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojdstradesstatisticsdiffAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPostStatus is PostStatus Setter
// 需要比较的状态
func (r *TaobaojdstradesstatisticsdiffAPIRequest) SetPostStatus(_postStatus string) error {
	r._postStatus = _postStatus
	r.Set("post_status", _postStatus)
	return nil
}

// GetPostStatus PostStatus Getter
func (r TaobaojdstradesstatisticsdiffAPIRequest) GetPostStatus() string {
	return r._postStatus
}

// SetPreStatus is PreStatus Setter
// 需要比较的状态，将会和post_status做比较
func (r *TaobaojdstradesstatisticsdiffAPIRequest) SetPreStatus(_preStatus string) error {
	r._preStatus = _preStatus
	r.Set("pre_status", _preStatus)
	return nil
}

// GetPreStatus PreStatus Getter
func (r TaobaojdstradesstatisticsdiffAPIRequest) GetPreStatus() string {
	return r._preStatus
}

// SetDate is Date Setter
// 查询的日期，格式如YYYYMMDD的日期对应的数字
func (r *TaobaojdstradesstatisticsdiffAPIRequest) SetDate(_date int64) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r TaobaojdstradesstatisticsdiffAPIRequest) GetDate() int64 {
	return r._date
}

// SetPageNo is PageNo Setter
// 页数
func (r *TaobaojdstradesstatisticsdiffAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaojdstradesstatisticsdiffAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
