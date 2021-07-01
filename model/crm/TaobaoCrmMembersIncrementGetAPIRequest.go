package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmMembersIncrementGetAPIRequest
增量获取卖家会员 API请求
taobao.crm.members.increment.get

增量获取会员列表，接口返回符合查询条件的所有会员。任何状态更改都会返回,最大允许100 */
type TaobaoCrmMembersIncrementGetAPIRequest struct {
	model.Params
	// 会员等级
	_grade int64
	// 卖家修改会员信息的时间起点.
	_startModify string
	// 卖家修改会员信息的时间终点.如果不填写此字段,默认为当前时间.
	_endModify string
	// 每页显示的会员数，page_size的值不能超过100，最小值要大于1
	_pageSize int64
	// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1
	_currentPage int64
}

// NewTaobaoCrmMembersIncrementGetRequest 初始化TaobaoCrmMembersIncrementGetAPIRequest对象
func NewTaobaoCrmMembersIncrementGetRequest() *TaobaoCrmMembersIncrementGetAPIRequest {
	return &TaobaoCrmMembersIncrementGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMembersIncrementGetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.members.increment.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMembersIncrementGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Grade Setter
// 会员等级
func (r *TaobaoCrmMembersIncrementGetAPIRequest) SetGrade(_grade int64) error {
	r._grade = _grade
	r.Set("grade", _grade)
	return nil
}

// Get Grade Getter
func (r TaobaoCrmMembersIncrementGetAPIRequest) GetGrade() int64 {
	return r._grade
}

// Set is StartModify Setter
// 卖家修改会员信息的时间起点.
func (r *TaobaoCrmMembersIncrementGetAPIRequest) SetStartModify(_startModify string) error {
	r._startModify = _startModify
	r.Set("start_modify", _startModify)
	return nil
}

// Get StartModify Getter
func (r TaobaoCrmMembersIncrementGetAPIRequest) GetStartModify() string {
	return r._startModify
}

// Set is EndModify Setter
// 卖家修改会员信息的时间终点.如果不填写此字段,默认为当前时间.
func (r *TaobaoCrmMembersIncrementGetAPIRequest) SetEndModify(_endModify string) error {
	r._endModify = _endModify
	r.Set("end_modify", _endModify)
	return nil
}

// Get EndModify Getter
func (r TaobaoCrmMembersIncrementGetAPIRequest) GetEndModify() string {
	return r._endModify
}

// Set is PageSize Setter
// 每页显示的会员数，page_size的值不能超过100，最小值要大于1
func (r *TaobaoCrmMembersIncrementGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoCrmMembersIncrementGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is CurrentPage Setter
// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1
func (r *TaobaoCrmMembersIncrementGetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// Get CurrentPage Getter
func (r TaobaoCrmMembersIncrementGetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
