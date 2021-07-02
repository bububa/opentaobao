package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlihealthDrugStoreSearchAPIRequest 药品店内搜索 API请求
// taobao.alihealth.drug.store.search
//
// 提供给千牛智能客服，在阿里健康O2O店铺内搜索药品
type TaobaoAlihealthDrugStoreSearchAPIRequest struct {
	model.Params
	// 搜索关键字
	_keyword string
	// 每页显示数量
	_pageSize int64
	// 店铺ID
	_shopId string
	// 页码
	_pageNo int64
}

// NewTaobaoAlihealthDrugStoreSearchRequest 初始化TaobaoAlihealthDrugStoreSearchAPIRequest对象
func NewTaobaoAlihealthDrugStoreSearchRequest() *TaobaoAlihealthDrugStoreSearchAPIRequest {
	return &TaobaoAlihealthDrugStoreSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlihealthDrugStoreSearchAPIRequest) GetApiMethodName() string {
	return "taobao.alihealth.drug.store.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlihealthDrugStoreSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetKeyword is Keyword Setter
// 搜索关键字
func (r *TaobaoAlihealthDrugStoreSearchAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r TaobaoAlihealthDrugStoreSearchAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetPageSize is PageSize Setter
// 每页显示数量
func (r *TaobaoAlihealthDrugStoreSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoAlihealthDrugStoreSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetShopId is ShopId Setter
// 店铺ID
func (r *TaobaoAlihealthDrugStoreSearchAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TaobaoAlihealthDrugStoreSearchAPIRequest) GetShopId() string {
	return r._shopId
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoAlihealthDrugStoreSearchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoAlihealthDrugStoreSearchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
