package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxmaterialmaterialfindpageAPIRequest 获取商品池 API请求
// taobao.onebp.dkx.material.material.findpage
//
// 获取商品池。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa。
type TaobaoonebpdkxmaterialmaterialfindpageAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 创意数据
	_materialQuery *MaterialQueryTopDto
}

// NewTaobaoonebpdkxmaterialmaterialfindpageRequest 初始化TaobaoonebpdkxmaterialmaterialfindpageAPIRequest对象
func NewTaobaoonebpdkxmaterialmaterialfindpageRequest() *TaobaoonebpdkxmaterialmaterialfindpageAPIRequest {
	return &TaobaoonebpdkxmaterialmaterialfindpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoonebpdkxmaterialmaterialfindpageAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.material.material.findpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoonebpdkxmaterialmaterialfindpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoonebpdkxmaterialmaterialfindpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoonebpdkxmaterialmaterialfindpageAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoonebpdkxmaterialmaterialfindpageAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetMaterialQuery is MaterialQuery Setter
// 创意数据
func (r *TaobaoonebpdkxmaterialmaterialfindpageAPIRequest) SetMaterialQuery(_materialQuery *MaterialQueryTopDto) error {
	r._materialQuery = _materialQuery
	r.Set("material_query", _materialQuery)
	return nil
}

// GetMaterialQuery MaterialQuery Getter
func (r TaobaoonebpdkxmaterialmaterialfindpageAPIRequest) GetMaterialQuery() *MaterialQueryTopDto {
	return r._materialQuery
}
