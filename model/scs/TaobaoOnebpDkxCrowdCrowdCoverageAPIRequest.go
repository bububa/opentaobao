package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxcrowdcrowdcoverageAPIRequest 获取人数预估 API请求
// taobao.onebp.dkx.crowd.crowd.coverage
//
// 获取人数预估，场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoonebpdkxcrowdcrowdcoverageAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 官方人群查询
	_groupQuery *GroupQueryTopDto
}

// NewTaobaoonebpdkxcrowdcrowdcoverageRequest 初始化TaobaoonebpdkxcrowdcrowdcoverageAPIRequest对象
func NewTaobaoonebpdkxcrowdcrowdcoverageRequest() *TaobaoonebpdkxcrowdcrowdcoverageAPIRequest {
	return &TaobaoonebpdkxcrowdcrowdcoverageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoonebpdkxcrowdcrowdcoverageAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.crowd.crowd.coverage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoonebpdkxcrowdcrowdcoverageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoonebpdkxcrowdcrowdcoverageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoonebpdkxcrowdcrowdcoverageAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoonebpdkxcrowdcrowdcoverageAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetGroupQuery is GroupQuery Setter
// 官方人群查询
func (r *TaobaoonebpdkxcrowdcrowdcoverageAPIRequest) SetGroupQuery(_groupQuery *GroupQueryTopDto) error {
	r._groupQuery = _groupQuery
	r.Set("group_query", _groupQuery)
	return nil
}

// GetGroupQuery GroupQuery Getter
func (r TaobaoonebpdkxcrowdcrowdcoverageAPIRequest) GetGroupQuery() *GroupQueryTopDto {
	return r._groupQuery
}
