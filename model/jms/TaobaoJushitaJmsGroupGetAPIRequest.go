package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojushitajmsgroupgetAPIRequest 查询ONS分组 API请求
// taobao.jushita.jms.group.get
//
// 查询当前appkey在ONS中已有的分组
type TaobaojushitajmsgroupgetAPIRequest struct {
	model.Params
	// 要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。
	_groupNames []string
	// 页码
	_pageNo int64
	// 每页返回多少个分组
	_pageSize int64
}

// NewTaobaojushitajmsgroupgetRequest 初始化TaobaojushitajmsgroupgetAPIRequest对象
func NewTaobaojushitajmsgroupgetRequest() *TaobaojushitajmsgroupgetAPIRequest {
	return &TaobaojushitajmsgroupgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojushitajmsgroupgetAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jms.group.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojushitajmsgroupgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojushitajmsgroupgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupNames is GroupNames Setter
// 要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。
func (r *TaobaojushitajmsgroupgetAPIRequest) SetGroupNames(_groupNames []string) error {
	r._groupNames = _groupNames
	r.Set("group_names", _groupNames)
	return nil
}

// GetGroupNames GroupNames Getter
func (r TaobaojushitajmsgroupgetAPIRequest) GetGroupNames() []string {
	return r._groupNames
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaojushitajmsgroupgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaojushitajmsgroupgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页返回多少个分组
func (r *TaobaojushitajmsgroupgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaojushitajmsgroupgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
