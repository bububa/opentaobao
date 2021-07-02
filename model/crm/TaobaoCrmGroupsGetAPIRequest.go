package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGroupsGetAPIRequest 查询卖家的分组 API请求
// taobao.crm.groups.get
//
// 查询卖家的分组，返回查询到的分组列表，分页返回分组
type TaobaoCrmGroupsGetAPIRequest struct {
	model.Params
	// 每页显示的记录数，其最大值不能超过100条，最小值为1，默认20条
	_pageSize int64
	// 显示第几页的分组，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页码为1
	_currentPage int64
}

// NewTaobaoCrmGroupsGetRequest 初始化TaobaoCrmGroupsGetAPIRequest对象
func NewTaobaoCrmGroupsGetRequest() *TaobaoCrmGroupsGetAPIRequest {
	return &TaobaoCrmGroupsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGroupsGetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.groups.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGroupsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPageSize is PageSize Setter
// 每页显示的记录数，其最大值不能超过100条，最小值为1，默认20条
func (r *TaobaoCrmGroupsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoCrmGroupsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 显示第几页的分组，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页码为1
func (r *TaobaoCrmGroupsGetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoCrmGroupsGetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
