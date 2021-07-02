package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest 销量明星api相关接口 API请求
// taobao.simba.salestar.keywords.recommend.get
//
// 取得一个推广组的推荐关键词列表
type TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest struct {
	model.Params
	// 推广组ID
	_adgroupId int64
	// 产品类型101001005代表标准推广，101001014代表销量明星
	_productId int64
}

// NewTaobaoSimbaSalestarKeywordsRecommendGetRequest 初始化TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest对象
func NewTaobaoSimbaSalestarKeywordsRecommendGetRequest() *TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest {
	return &TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.keywords.recommend.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAdgroupId is AdgroupId Setter
// 推广组ID
func (r *TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetProductId is ProductId Setter
// 产品类型101001005代表标准推广，101001014代表销量明星
func (r *TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest) GetProductId() int64 {
	return r._productId
}
