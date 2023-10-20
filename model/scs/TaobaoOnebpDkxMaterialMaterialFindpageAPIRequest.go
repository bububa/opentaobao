package scs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest 获取商品池 API请求
// taobao.onebp.dkx.material.material.findpage
//
// 获取商品池。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa。
type TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 创意数据
	_materialQuery *MaterialQueryTopDto
}

// NewTaobaoOnebpDkxMaterialMaterialFindpageRequest 初始化TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest对象
func NewTaobaoOnebpDkxMaterialMaterialFindpageRequest() *TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest {
	return &TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest) Reset() {
	r._apiServiceContext = nil
	r._materialQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.material.material.findpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetMaterialQuery is MaterialQuery Setter
// 创意数据
func (r *TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest) SetMaterialQuery(_materialQuery *MaterialQueryTopDto) error {
	r._materialQuery = _materialQuery
	r.Set("material_query", _materialQuery)
	return nil
}

// GetMaterialQuery MaterialQuery Getter
func (r TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest) GetMaterialQuery() *MaterialQueryTopDto {
	return r._materialQuery
}

var poolTaobaoOnebpDkxMaterialMaterialFindpageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOnebpDkxMaterialMaterialFindpageRequest()
	},
}

// GetTaobaoOnebpDkxMaterialMaterialFindpageRequest 从 sync.Pool 获取 TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest
func GetTaobaoOnebpDkxMaterialMaterialFindpageAPIRequest() *TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest {
	return poolTaobaoOnebpDkxMaterialMaterialFindpageAPIRequest.Get().(*TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest)
}

// ReleaseTaobaoOnebpDkxMaterialMaterialFindpageAPIRequest 将 TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest 放入 sync.Pool
func ReleaseTaobaoOnebpDkxMaterialMaterialFindpageAPIRequest(v *TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest) {
	v.Reset()
	poolTaobaoOnebpDkxMaterialMaterialFindpageAPIRequest.Put(v)
}
