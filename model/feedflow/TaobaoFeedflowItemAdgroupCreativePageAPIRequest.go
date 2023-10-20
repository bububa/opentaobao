package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgroupcreativepageAPIRequest 信息流单元下查看创意 API请求
// taobao.feedflow.item.adgroup.creative.page
//
// 信息流单元下查看创意
type TaobaofeedflowitemadgroupcreativepageAPIRequest struct {
	model.Params
	// 绑定查询条件
	_creativeBindQuery *CreativeBindQueryDto
}

// NewTaobaofeedflowitemadgroupcreativepageRequest 初始化TaobaofeedflowitemadgroupcreativepageAPIRequest对象
func NewTaobaofeedflowitemadgroupcreativepageRequest() *TaobaofeedflowitemadgroupcreativepageAPIRequest {
	return &TaobaofeedflowitemadgroupcreativepageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemadgroupcreativepageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.creative.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemadgroupcreativepageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemadgroupcreativepageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreativeBindQuery is CreativeBindQuery Setter
// 绑定查询条件
func (r *TaobaofeedflowitemadgroupcreativepageAPIRequest) SetCreativeBindQuery(_creativeBindQuery *CreativeBindQueryDto) error {
	r._creativeBindQuery = _creativeBindQuery
	r.Set("creative_bind_query", _creativeBindQuery)
	return nil
}

// GetCreativeBindQuery CreativeBindQuery Getter
func (r TaobaofeedflowitemadgroupcreativepageAPIRequest) GetCreativeBindQuery() *CreativeBindQueryDto {
	return r._creativeBindQuery
}
