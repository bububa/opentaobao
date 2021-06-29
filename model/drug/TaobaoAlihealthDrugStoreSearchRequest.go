package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
药品店内搜索 API请求
taobao.alihealth.drug.store.search

提供给千牛智能客服，在阿里健康O2O店铺内搜索药品
*/
type TaobaoAlihealthDrugStoreSearchRequest struct {
    model.Params
    // 搜索关键字
    _keyword   string
    // 每页显示数量
    _pageSize   int64
    // 店铺ID
    _shopId   string
    // 页码
    _pageNo   int64
}

// 初始化TaobaoAlihealthDrugStoreSearchRequest对象
func NewTaobaoAlihealthDrugStoreSearchRequest() *TaobaoAlihealthDrugStoreSearchRequest{
    return &TaobaoAlihealthDrugStoreSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlihealthDrugStoreSearchRequest) GetApiMethodName() string {
    return "taobao.alihealth.drug.store.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlihealthDrugStoreSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Keyword Setter
// 搜索关键字
func (r *TaobaoAlihealthDrugStoreSearchRequest) SetKeyword(_keyword string) error {
    r._keyword = _keyword
    r.Set("keyword", _keyword)
    return nil
}

// Keyword Getter
func (r TaobaoAlihealthDrugStoreSearchRequest) GetKeyword() string {
    return r._keyword
}
// PageSize Setter
// 每页显示数量
func (r *TaobaoAlihealthDrugStoreSearchRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoAlihealthDrugStoreSearchRequest) GetPageSize() int64 {
    return r._pageSize
}
// ShopId Setter
// 店铺ID
func (r *TaobaoAlihealthDrugStoreSearchRequest) SetShopId(_shopId string) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r TaobaoAlihealthDrugStoreSearchRequest) GetShopId() string {
    return r._shopId
}
// PageNo Setter
// 页码
func (r *TaobaoAlihealthDrugStoreSearchRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoAlihealthDrugStoreSearchRequest) GetPageNo() int64 {
    return r._pageNo
}
