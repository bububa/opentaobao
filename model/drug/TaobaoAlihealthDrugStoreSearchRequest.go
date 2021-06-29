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
    keyword   string
    // 每页显示数量
    pageSize   int64
    // 店铺ID
    shopId   string
    // 页码
    pageNo   int64
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
func (r *TaobaoAlihealthDrugStoreSearchRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

// Keyword Getter
func (r TaobaoAlihealthDrugStoreSearchRequest) GetKeyword() string {
    return r.keyword
}
// PageSize Setter
// 每页显示数量
func (r *TaobaoAlihealthDrugStoreSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoAlihealthDrugStoreSearchRequest) GetPageSize() int64 {
    return r.pageSize
}
// ShopId Setter
// 店铺ID
func (r *TaobaoAlihealthDrugStoreSearchRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r TaobaoAlihealthDrugStoreSearchRequest) GetShopId() string {
    return r.shopId
}
// PageNo Setter
// 页码
func (r *TaobaoAlihealthDrugStoreSearchRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoAlihealthDrugStoreSearchRequest) GetPageNo() int64 {
    return r.pageNo
}
