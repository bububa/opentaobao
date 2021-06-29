package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家外卖店列表 APIRequest
taobao.drug.shop.list

查询卖家外卖店列表
*/
type TaobaoDrugShopListRequest struct {
    model.Params

    // 查询关键字
    keywords   string 

    // 店铺状态，歇业：0，营业：1，所有：-1
    status   int64 

    // 页码
    page   int64 

    // 每页条数
    pageSize   int64 

}

func NewTaobaoDrugShopListRequest() *TaobaoDrugShopListRequest{
    return &TaobaoDrugShopListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDrugShopListRequest) GetApiMethodName() string {
    return "taobao.drug.shop.list"
}

func (r TaobaoDrugShopListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDrugShopListRequest) SetKeywords(keywords string) error {
    r.keywords = keywords
    r.Set("keywords", keywords)
    return nil
}

func (r TaobaoDrugShopListRequest) GetKeywords() string {
    return r.keywords
}

func (r *TaobaoDrugShopListRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoDrugShopListRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoDrugShopListRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r TaobaoDrugShopListRequest) GetPage() int64 {
    return r.page
}

func (r *TaobaoDrugShopListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoDrugShopListRequest) GetPageSize() int64 {
    return r.pageSize
}

