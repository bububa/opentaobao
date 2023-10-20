package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcrowdpageAPIRequest 分页查询单品单元下人群列表 API请求
// taobao.feedflow.item.crowd.page
//
// 分页查询单品单元下人群列表
type TaobaofeedflowitemcrowdpageAPIRequest struct {
	model.Params
	// 查询条件
	_crowdQuery *CrowdQueryDto
}

// NewTaobaofeedflowitemcrowdpageRequest 初始化TaobaofeedflowitemcrowdpageAPIRequest对象
func NewTaobaofeedflowitemcrowdpageRequest() *TaobaofeedflowitemcrowdpageAPIRequest {
	return &TaobaofeedflowitemcrowdpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcrowdpageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcrowdpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcrowdpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrowdQuery is CrowdQuery Setter
// 查询条件
func (r *TaobaofeedflowitemcrowdpageAPIRequest) SetCrowdQuery(_crowdQuery *CrowdQueryDto) error {
	r._crowdQuery = _crowdQuery
	r.Set("crowd_query", _crowdQuery)
	return nil
}

// GetCrowdQuery CrowdQuery Getter
func (r TaobaofeedflowitemcrowdpageAPIRequest) GetCrowdQuery() *CrowdQueryDto {
	return r._crowdQuery
}
