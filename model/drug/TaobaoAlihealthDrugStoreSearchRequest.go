package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
药品店内搜索 APIRequest
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

func NewTaobaoAlihealthDrugStoreSearchRequest() *TaobaoAlihealthDrugStoreSearchRequest{
    return &TaobaoAlihealthDrugStoreSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlihealthDrugStoreSearchRequest) GetApiMethodName() string {
    return "taobao.alihealth.drug.store.search"
}

func (r TaobaoAlihealthDrugStoreSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlihealthDrugStoreSearchRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

func (r TaobaoAlihealthDrugStoreSearchRequest) GetKeyword() string {
    return r.keyword
}

func (r *TaobaoAlihealthDrugStoreSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoAlihealthDrugStoreSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoAlihealthDrugStoreSearchRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r TaobaoAlihealthDrugStoreSearchRequest) GetShopId() string {
    return r.shopId
}

func (r *TaobaoAlihealthDrugStoreSearchRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoAlihealthDrugStoreSearchRequest) GetPageNo() int64 {
    return r.pageNo
}

