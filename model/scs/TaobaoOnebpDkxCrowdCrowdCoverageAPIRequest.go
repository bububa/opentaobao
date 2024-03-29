package scs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest 获取人数预估 API请求
// taobao.onebp.dkx.crowd.crowd.coverage
//
// 获取人数预估，场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 官方人群查询
	_groupQuery *GroupQueryTopDto
}

// NewTaobaoOnebpDkxCrowdCrowdCoverageRequest 初始化TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest对象
func NewTaobaoOnebpDkxCrowdCrowdCoverageRequest() *TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest {
	return &TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest) Reset() {
	r._apiServiceContext = nil
	r._groupQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.crowd.crowd.coverage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetGroupQuery is GroupQuery Setter
// 官方人群查询
func (r *TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest) SetGroupQuery(_groupQuery *GroupQueryTopDto) error {
	r._groupQuery = _groupQuery
	r.Set("group_query", _groupQuery)
	return nil
}

// GetGroupQuery GroupQuery Getter
func (r TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest) GetGroupQuery() *GroupQueryTopDto {
	return r._groupQuery
}

var poolTaobaoOnebpDkxCrowdCrowdCoverageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOnebpDkxCrowdCrowdCoverageRequest()
	},
}

// GetTaobaoOnebpDkxCrowdCrowdCoverageRequest 从 sync.Pool 获取 TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest
func GetTaobaoOnebpDkxCrowdCrowdCoverageAPIRequest() *TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest {
	return poolTaobaoOnebpDkxCrowdCrowdCoverageAPIRequest.Get().(*TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest)
}

// ReleaseTaobaoOnebpDkxCrowdCrowdCoverageAPIRequest 将 TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest 放入 sync.Pool
func ReleaseTaobaoOnebpDkxCrowdCrowdCoverageAPIRequest(v *TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest) {
	v.Reset()
	poolTaobaoOnebpDkxCrowdCrowdCoverageAPIRequest.Put(v)
}
