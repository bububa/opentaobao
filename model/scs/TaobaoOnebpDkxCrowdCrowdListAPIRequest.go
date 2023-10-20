package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCrowdCrowdListAPIRequest 获取人群信息列表 API请求
// taobao.onebp.dkx.crowd.crowd.list
//
// 获取人群信息列表。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCrowdCrowdListAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 官方人群查询
	_groupQuery *GroupQueryTopDto
}

// NewTaobaoOnebpDkxCrowdCrowdListRequest 初始化TaobaoOnebpDkxCrowdCrowdListAPIRequest对象
func NewTaobaoOnebpDkxCrowdCrowdListRequest() *TaobaoOnebpDkxCrowdCrowdListAPIRequest {
	return &TaobaoOnebpDkxCrowdCrowdListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxCrowdCrowdListAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.crowd.crowd.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxCrowdCrowdListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOnebpDkxCrowdCrowdListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxCrowdCrowdListAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxCrowdCrowdListAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetGroupQuery is GroupQuery Setter
// 官方人群查询
func (r *TaobaoOnebpDkxCrowdCrowdListAPIRequest) SetGroupQuery(_groupQuery *GroupQueryTopDto) error {
	r._groupQuery = _groupQuery
	r.Set("group_query", _groupQuery)
	return nil
}

// GetGroupQuery GroupQuery Getter
func (r TaobaoOnebpDkxCrowdCrowdListAPIRequest) GetGroupQuery() *GroupQueryTopDto {
	return r._groupQuery
}
