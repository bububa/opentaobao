package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxcrowdcrowdfindcrowdgroupsAPIRequest 获取人群组 API请求
// taobao.onebp.dkx.crowd.crowd.findcrowdgroups
//
// 获取人群组
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{ &#34;market_scene&#34;: &#34;ad_strategy_laxin&#34;}
type TaobaoonebpdkxcrowdcrowdfindcrowdgroupsAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 官方人群查询
	_groupQuery *GroupQueryTopDto
}

// NewTaobaoonebpdkxcrowdcrowdfindcrowdgroupsRequest 初始化TaobaoonebpdkxcrowdcrowdfindcrowdgroupsAPIRequest对象
func NewTaobaoonebpdkxcrowdcrowdfindcrowdgroupsRequest() *TaobaoonebpdkxcrowdcrowdfindcrowdgroupsAPIRequest {
	return &TaobaoonebpdkxcrowdcrowdfindcrowdgroupsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoonebpdkxcrowdcrowdfindcrowdgroupsAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.crowd.crowd.findcrowdgroups"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoonebpdkxcrowdcrowdfindcrowdgroupsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoonebpdkxcrowdcrowdfindcrowdgroupsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoonebpdkxcrowdcrowdfindcrowdgroupsAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoonebpdkxcrowdcrowdfindcrowdgroupsAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetGroupQuery is GroupQuery Setter
// 官方人群查询
func (r *TaobaoonebpdkxcrowdcrowdfindcrowdgroupsAPIRequest) SetGroupQuery(_groupQuery *GroupQueryTopDto) error {
	r._groupQuery = _groupQuery
	r.Set("group_query", _groupQuery)
	return nil
}

// GetGroupQuery GroupQuery Getter
func (r TaobaoonebpdkxcrowdcrowdfindcrowdgroupsAPIRequest) GetGroupQuery() *GroupQueryTopDto {
	return r._groupQuery
}
