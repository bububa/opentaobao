package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxcrowdcrowdtemplateAPIRequest 获取人群模版 API请求
// taobao.onebp.dkx.crowd.crowd.template
//
// 获取人群模版，场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoonebpdkxcrowdcrowdtemplateAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 官方人群查询
	_groupQuery *GroupQueryTopDto
}

// NewTaobaoonebpdkxcrowdcrowdtemplateRequest 初始化TaobaoonebpdkxcrowdcrowdtemplateAPIRequest对象
func NewTaobaoonebpdkxcrowdcrowdtemplateRequest() *TaobaoonebpdkxcrowdcrowdtemplateAPIRequest {
	return &TaobaoonebpdkxcrowdcrowdtemplateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoonebpdkxcrowdcrowdtemplateAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.crowd.crowd.template"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoonebpdkxcrowdcrowdtemplateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoonebpdkxcrowdcrowdtemplateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoonebpdkxcrowdcrowdtemplateAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoonebpdkxcrowdcrowdtemplateAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetGroupQuery is GroupQuery Setter
// 官方人群查询
func (r *TaobaoonebpdkxcrowdcrowdtemplateAPIRequest) SetGroupQuery(_groupQuery *GroupQueryTopDto) error {
	r._groupQuery = _groupQuery
	r.Set("group_query", _groupQuery)
	return nil
}

// GetGroupQuery GroupQuery Getter
func (r TaobaoonebpdkxcrowdcrowdtemplateAPIRequest) GetGroupQuery() *GroupQueryTopDto {
	return r._groupQuery
}
