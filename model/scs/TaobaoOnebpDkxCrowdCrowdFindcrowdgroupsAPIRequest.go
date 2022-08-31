package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIRequest 获取人群组 API请求
// taobao.onebp.dkx.crowd.crowd.findcrowdgroups
//
// 获取人群组
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{ &#34;market_scene&#34;: &#34;ad_strategy_laxin&#34;}
type TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 官方人群查询
	_groupQuery *GroupQueryTopDto
}

// NewTaobaoOnebpDkxCrowdCrowdFindcrowdgroupsRequest 初始化TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIRequest对象
func NewTaobaoOnebpDkxCrowdCrowdFindcrowdgroupsRequest() *TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIRequest {
	return &TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.crowd.crowd.findcrowdgroups"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetGroupQuery is GroupQuery Setter
// 官方人群查询
func (r *TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIRequest) SetGroupQuery(_groupQuery *GroupQueryTopDto) error {
	r._groupQuery = _groupQuery
	r.Set("group_query", _groupQuery)
	return nil
}

// GetGroupQuery GroupQuery Getter
func (r TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIRequest) GetGroupQuery() *GroupQueryTopDto {
	return r._groupQuery
}
