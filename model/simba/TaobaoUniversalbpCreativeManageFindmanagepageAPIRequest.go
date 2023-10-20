package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpcreativemanagefindmanagepageAPIRequest 创意库查询创意列表 API请求
// taobao.universalbp.creative.manage.findmanagepage
//
// 创意库查询创意列表
type TaobaouniversalbpcreativemanagefindmanagepageAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// creativeQueryVO
	_creativeQueryVO *CreativeQueryVo
}

// NewTaobaouniversalbpcreativemanagefindmanagepageRequest 初始化TaobaouniversalbpcreativemanagefindmanagepageAPIRequest对象
func NewTaobaouniversalbpcreativemanagefindmanagepageRequest() *TaobaouniversalbpcreativemanagefindmanagepageAPIRequest {
	return &TaobaouniversalbpcreativemanagefindmanagepageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpcreativemanagefindmanagepageAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.creative.manage.findmanagepage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpcreativemanagefindmanagepageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpcreativemanagefindmanagepageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpcreativemanagefindmanagepageAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpcreativemanagefindmanagepageAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetCreativeQueryVO is CreativeQueryVO Setter
// creativeQueryVO
func (r *TaobaouniversalbpcreativemanagefindmanagepageAPIRequest) SetCreativeQueryVO(_creativeQueryVO *CreativeQueryVo) error {
	r._creativeQueryVO = _creativeQueryVO
	r.Set("creative_query_v_o", _creativeQueryVO)
	return nil
}

// GetCreativeQueryVO CreativeQueryVO Getter
func (r TaobaouniversalbpcreativemanagefindmanagepageAPIRequest) GetCreativeQueryVO() *CreativeQueryVo {
	return r._creativeQueryVO
}
