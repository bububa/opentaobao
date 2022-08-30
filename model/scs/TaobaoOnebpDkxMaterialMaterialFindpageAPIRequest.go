package scs

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.material.material.findpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
