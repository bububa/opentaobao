package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmmembersincrementgetAPIRequest 增量获取卖家会员 API请求
// taobao.crm.members.increment.get
//
// 增量获取会员列表，接口返回符合查询条件的所有会员。任何状态更改都会返回,最大允许100
type TaobaocrmmembersincrementgetAPIRequest struct {
	model.Params
	// 卖家修改会员信息的时间起点.
	_startModify string
	// 卖家修改会员信息的时间终点.如果不填写此字段,默认为当前时间.
	_endModify string
	// 每页显示的会员数，page_size的值不能超过100，最小值要大于1
	_pageSize int64
	// 会员等级
	_grade int64
	// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1
	_currentPage int64
}

// NewTaobaocrmmembersincrementgetRequest 初始化TaobaocrmmembersincrementgetAPIRequest对象
func NewTaobaocrmmembersincrementgetRequest() *TaobaocrmmembersincrementgetAPIRequest {
	return &TaobaocrmmembersincrementgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmmembersincrementgetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.members.increment.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmmembersincrementgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmmembersincrementgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartModify is StartModify Setter
// 卖家修改会员信息的时间起点.
func (r *TaobaocrmmembersincrementgetAPIRequest) SetStartModify(_startModify string) error {
	r._startModify = _startModify
	r.Set("start_modify", _startModify)
	return nil
}

// GetStartModify StartModify Getter
func (r TaobaocrmmembersincrementgetAPIRequest) GetStartModify() string {
	return r._startModify
}

// SetEndModify is EndModify Setter
// 卖家修改会员信息的时间终点.如果不填写此字段,默认为当前时间.
func (r *TaobaocrmmembersincrementgetAPIRequest) SetEndModify(_endModify string) error {
	r._endModify = _endModify
	r.Set("end_modify", _endModify)
	return nil
}

// GetEndModify EndModify Getter
func (r TaobaocrmmembersincrementgetAPIRequest) GetEndModify() string {
	return r._endModify
}

// SetPageSize is PageSize Setter
// 每页显示的会员数，page_size的值不能超过100，最小值要大于1
func (r *TaobaocrmmembersincrementgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaocrmmembersincrementgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetGrade is Grade Setter
// 会员等级
func (r *TaobaocrmmembersincrementgetAPIRequest) SetGrade(_grade int64) error {
	r._grade = _grade
	r.Set("grade", _grade)
	return nil
}

// GetGrade Grade Getter
func (r TaobaocrmmembersincrementgetAPIRequest) GetGrade() int64 {
	return r._grade
}

// SetCurrentPage is CurrentPage Setter
// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1
func (r *TaobaocrmmembersincrementgetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaocrmmembersincrementgetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
