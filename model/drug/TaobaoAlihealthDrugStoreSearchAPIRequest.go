package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalihealthdrugstoresearchAPIRequest 药品店内搜索 API请求
// taobao.alihealth.drug.store.search
//
// 提供给千牛智能客服，在阿里健康O2O店铺内搜索药品
type TaobaoalihealthdrugstoresearchAPIRequest struct {
	model.Params
	// 搜索关键字
	_keyword string
	// 店铺ID
	_shopId string
	// 每页显示数量
	_pageSize int64
	// 页码
	_pageNo int64
}

// NewTaobaoalihealthdrugstoresearchRequest 初始化TaobaoalihealthdrugstoresearchAPIRequest对象
func NewTaobaoalihealthdrugstoresearchRequest() *TaobaoalihealthdrugstoresearchAPIRequest {
	return &TaobaoalihealthdrugstoresearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalihealthdrugstoresearchAPIRequest) GetApiMethodName() string {
	return "taobao.alihealth.drug.store.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalihealthdrugstoresearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalihealthdrugstoresearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 搜索关键字
func (r *TaobaoalihealthdrugstoresearchAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r TaobaoalihealthdrugstoresearchAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetShopId is ShopId Setter
// 店铺ID
func (r *TaobaoalihealthdrugstoresearchAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TaobaoalihealthdrugstoresearchAPIRequest) GetShopId() string {
	return r._shopId
}

// SetPageSize is PageSize Setter
// 每页显示数量
func (r *TaobaoalihealthdrugstoresearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoalihealthdrugstoresearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoalihealthdrugstoresearchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoalihealthdrugstoresearchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
