package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest 查询推荐人群 API请求
// taobao.universalbp.crowd.findrecommendcrowd
//
// 入参推广信息，查询推荐人群，查出的推荐人群列表，不需转换，直接可用于入参其他计划创编类接口
type TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// crowdRecQueryVO
	_crowdRecQueryVO *CrowdRecQueryVo
}

// NewTaobaoUniversalbpCrowdFindrecommendcrowdRequest 初始化TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest对象
func NewTaobaoUniversalbpCrowdFindrecommendcrowdRequest() *TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest {
	return &TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.crowd.findrecommendcrowd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetCrowdRecQueryVO is CrowdRecQueryVO Setter
// crowdRecQueryVO
func (r *TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest) SetCrowdRecQueryVO(_crowdRecQueryVO *CrowdRecQueryVo) error {
	r._crowdRecQueryVO = _crowdRecQueryVO
	r.Set("crowd_rec_query_v_o", _crowdRecQueryVO)
	return nil
}

// GetCrowdRecQueryVO CrowdRecQueryVO Getter
func (r TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest) GetCrowdRecQueryVO() *CrowdRecQueryVo {
	return r._crowdRecQueryVO
}
