package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFuwuScoresGetAPIRequest 服务平台评价查询接口 API请求
// taobao.fuwu.scores.get
//
// 根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟
type TaobaoFuwuScoresGetAPIRequest struct {
	model.Params
	// 评价日期，查询某一天的评价
	_date string
	// 当前页
	_currentPage int64
	// 每页获取条数。默认值40，最小值1，最大值100。
	_pageSize int64
}

// NewTaobaoFuwuScoresGetRequest 初始化TaobaoFuwuScoresGetAPIRequest对象
func NewTaobaoFuwuScoresGetRequest() *TaobaoFuwuScoresGetAPIRequest {
	return &TaobaoFuwuScoresGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFuwuScoresGetAPIRequest) GetApiMethodName() string {
	return "taobao.fuwu.scores.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFuwuScoresGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDate is Date Setter
// 评价日期，查询某一天的评价
func (r *TaobaoFuwuScoresGetAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r TaobaoFuwuScoresGetAPIRequest) GetDate() string {
	return r._date
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *TaobaoFuwuScoresGetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoFuwuScoresGetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页获取条数。默认值40，最小值1，最大值100。
func (r *TaobaoFuwuScoresGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoFuwuScoresGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
