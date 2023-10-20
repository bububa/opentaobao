package scs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest 获取人群模版 API请求
// taobao.onebp.dkx.crowd.crowd.template
//
// 获取人群模版，场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 官方人群查询
	_groupQuery *GroupQueryTopDto
}

// NewTaobaoOnebpDkxCrowdCrowdTemplateRequest 初始化TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest对象
func NewTaobaoOnebpDkxCrowdCrowdTemplateRequest() *TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest {
	return &TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest) Reset() {
	r._apiServiceContext = nil
	r._groupQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.crowd.crowd.template"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetGroupQuery is GroupQuery Setter
// 官方人群查询
func (r *TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest) SetGroupQuery(_groupQuery *GroupQueryTopDto) error {
	r._groupQuery = _groupQuery
	r.Set("group_query", _groupQuery)
	return nil
}

// GetGroupQuery GroupQuery Getter
func (r TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest) GetGroupQuery() *GroupQueryTopDto {
	return r._groupQuery
}

var poolTaobaoOnebpDkxCrowdCrowdTemplateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOnebpDkxCrowdCrowdTemplateRequest()
	},
}

// GetTaobaoOnebpDkxCrowdCrowdTemplateRequest 从 sync.Pool 获取 TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest
func GetTaobaoOnebpDkxCrowdCrowdTemplateAPIRequest() *TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest {
	return poolTaobaoOnebpDkxCrowdCrowdTemplateAPIRequest.Get().(*TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest)
}

// ReleaseTaobaoOnebpDkxCrowdCrowdTemplateAPIRequest 将 TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest 放入 sync.Pool
func ReleaseTaobaoOnebpDkxCrowdCrowdTemplateAPIRequest(v *TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest) {
	v.Reset()
	poolTaobaoOnebpDkxCrowdCrowdTemplateAPIRequest.Put(v)
}
