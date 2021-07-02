package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcGroupsGetAPIRequest 获取自定义用户分组列表 API请求
// taobao.tmc.groups.get
//
// 获取自定义用户分组列表
type TaobaoTmcGroupsGetAPIRequest struct {
	model.Params
	// 要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。
	_groupNames []string
	// 页码
	_pageNo int64
	// 每页返回多少个分组
	_pageSize int64
}

// NewTaobaoTmcGroupsGetRequest 初始化TaobaoTmcGroupsGetAPIRequest对象
func NewTaobaoTmcGroupsGetRequest() *TaobaoTmcGroupsGetAPIRequest {
	return &TaobaoTmcGroupsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcGroupsGetAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.groups.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcGroupsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GroupNames Setter
// 要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。
func (r *TaobaoTmcGroupsGetAPIRequest) SetGroupNames(_groupNames []string) error {
	r._groupNames = _groupNames
	r.Set("group_names", _groupNames)
	return nil
}

// Get GroupNames Getter
func (r TaobaoTmcGroupsGetAPIRequest) GetGroupNames() []string {
	return r._groupNames
}

// Set is PageNo Setter
// 页码
func (r *TaobaoTmcGroupsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoTmcGroupsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 每页返回多少个分组
func (r *TaobaoTmcGroupsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoTmcGroupsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
