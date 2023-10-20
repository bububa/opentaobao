package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgroupadzonepageAPIRequest 信息流单元下查看绑定资源位 API请求
// taobao.feedflow.item.adgroup.adzone.page
//
// 信息流单元下查看绑定资源位
type TaobaofeedflowitemadgroupadzonepageAPIRequest struct {
	model.Params
	// 查询条件
	_adzoneBindQuery *AdzoneBindQueryDto
}

// NewTaobaofeedflowitemadgroupadzonepageRequest 初始化TaobaofeedflowitemadgroupadzonepageAPIRequest对象
func NewTaobaofeedflowitemadgroupadzonepageRequest() *TaobaofeedflowitemadgroupadzonepageAPIRequest {
	return &TaobaofeedflowitemadgroupadzonepageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemadgroupadzonepageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.adzone.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemadgroupadzonepageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemadgroupadzonepageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdzoneBindQuery is AdzoneBindQuery Setter
// 查询条件
func (r *TaobaofeedflowitemadgroupadzonepageAPIRequest) SetAdzoneBindQuery(_adzoneBindQuery *AdzoneBindQueryDto) error {
	r._adzoneBindQuery = _adzoneBindQuery
	r.Set("adzone_bind_query", _adzoneBindQuery)
	return nil
}

// GetAdzoneBindQuery AdzoneBindQuery Getter
func (r TaobaofeedflowitemadgroupadzonepageAPIRequest) GetAdzoneBindQuery() *AdzoneBindQueryDto {
	return r._adzoneBindQuery
}
