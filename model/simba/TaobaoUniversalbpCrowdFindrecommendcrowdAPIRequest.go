package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpcrowdfindrecommendcrowdAPIRequest 查询推荐人群 API请求
// taobao.universalbp.crowd.findrecommendcrowd
//
// 入参推广信息，查询推荐人群，查出的推荐人群列表，不需转换，直接可用于入参其他计划创编类接口
type TaobaouniversalbpcrowdfindrecommendcrowdAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// crowdRecQueryVO
	_crowdRecQueryVO *CrowdRecQueryVo
}

// NewTaobaouniversalbpcrowdfindrecommendcrowdRequest 初始化TaobaouniversalbpcrowdfindrecommendcrowdAPIRequest对象
func NewTaobaouniversalbpcrowdfindrecommendcrowdRequest() *TaobaouniversalbpcrowdfindrecommendcrowdAPIRequest {
	return &TaobaouniversalbpcrowdfindrecommendcrowdAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpcrowdfindrecommendcrowdAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.crowd.findrecommendcrowd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpcrowdfindrecommendcrowdAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpcrowdfindrecommendcrowdAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpcrowdfindrecommendcrowdAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpcrowdfindrecommendcrowdAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetCrowdRecQueryVO is CrowdRecQueryVO Setter
// crowdRecQueryVO
func (r *TaobaouniversalbpcrowdfindrecommendcrowdAPIRequest) SetCrowdRecQueryVO(_crowdRecQueryVO *CrowdRecQueryVo) error {
	r._crowdRecQueryVO = _crowdRecQueryVO
	r.Set("crowd_rec_query_v_o", _crowdRecQueryVO)
	return nil
}

// GetCrowdRecQueryVO CrowdRecQueryVO Getter
func (r TaobaouniversalbpcrowdfindrecommendcrowdAPIRequest) GetCrowdRecQueryVO() *CrowdRecQueryVo {
	return r._crowdRecQueryVO
}
