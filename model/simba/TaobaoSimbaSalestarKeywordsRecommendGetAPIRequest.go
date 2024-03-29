package simba

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest) Reset() {
	r._adgroupId = 0
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.keywords.recommend.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoSimbaSalestarKeywordsRecommendGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaSalestarKeywordsRecommendGetRequest()
	},
}

// GetTaobaoSimbaSalestarKeywordsRecommendGetRequest 从 sync.Pool 获取 TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest
func GetTaobaoSimbaSalestarKeywordsRecommendGetAPIRequest() *TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest {
	return poolTaobaoSimbaSalestarKeywordsRecommendGetAPIRequest.Get().(*TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest)
}

// ReleaseTaobaoSimbaSalestarKeywordsRecommendGetAPIRequest 将 TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaSalestarKeywordsRecommendGetAPIRequest(v *TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaSalestarKeywordsRecommendGetAPIRequest.Put(v)
}
