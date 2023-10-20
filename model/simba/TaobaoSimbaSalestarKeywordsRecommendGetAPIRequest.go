package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasalestarkeywordsrecommendgetAPIRequest 销量明星api相关接口 API请求
// taobao.simba.salestar.keywords.recommend.get
//
// 取得一个推广组的推荐关键词列表
type TaobaosimbasalestarkeywordsrecommendgetAPIRequest struct {
	model.Params
	// 推广组ID
	_adgroupId int64
	// 产品类型101001005代表标准推广，101001014代表销量明星
	_productId int64
}

// NewTaobaosimbasalestarkeywordsrecommendgetRequest 初始化TaobaosimbasalestarkeywordsrecommendgetAPIRequest对象
func NewTaobaosimbasalestarkeywordsrecommendgetRequest() *TaobaosimbasalestarkeywordsrecommendgetAPIRequest {
	return &TaobaosimbasalestarkeywordsrecommendgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbasalestarkeywordsrecommendgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.keywords.recommend.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbasalestarkeywordsrecommendgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbasalestarkeywordsrecommendgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupId is AdgroupId Setter
// 推广组ID
func (r *TaobaosimbasalestarkeywordsrecommendgetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbasalestarkeywordsrecommendgetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetProductId is ProductId Setter
// 产品类型101001005代表标准推广，101001014代表销量明星
func (r *TaobaosimbasalestarkeywordsrecommendgetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaosimbasalestarkeywordsrecommendgetAPIRequest) GetProductId() int64 {
	return r._productId
}
