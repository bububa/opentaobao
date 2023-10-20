package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgrouppageAPIRequest 查询单元列表 API请求
// taobao.feedflow.item.adgroup.page
//
// 通过计划id查询单元信息
type TaobaofeedflowitemadgrouppageAPIRequest struct {
	model.Params
	// 系统自动生成
	_adgroupQuery *AdgroupQueryDto
}

// NewTaobaofeedflowitemadgrouppageRequest 初始化TaobaofeedflowitemadgrouppageAPIRequest对象
func NewTaobaofeedflowitemadgrouppageRequest() *TaobaofeedflowitemadgrouppageAPIRequest {
	return &TaobaofeedflowitemadgrouppageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemadgrouppageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemadgrouppageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemadgrouppageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupQuery is AdgroupQuery Setter
// 系统自动生成
func (r *TaobaofeedflowitemadgrouppageAPIRequest) SetAdgroupQuery(_adgroupQuery *AdgroupQueryDto) error {
	r._adgroupQuery = _adgroupQuery
	r.Set("adgroup_query", _adgroupQuery)
	return nil
}

// GetAdgroupQuery AdgroupQuery Getter
func (r TaobaofeedflowitemadgrouppageAPIRequest) GetAdgroupQuery() *AdgroupQueryDto {
	return r._adgroupQuery
}
