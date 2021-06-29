package c2m

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺商品接口 APIRequest
taobao.txp.item.itemlistget

淘小铺商品的查询服务。
*/
type TaobaoTxpItemItemlistgetRequest struct {
    model.Params

    // 第几页
    beginPage   int64 

    // 每页多少条
    pageSize   int64 

}

func NewTaobaoTxpItemItemlistgetRequest() *TaobaoTxpItemItemlistgetRequest{
    return &TaobaoTxpItemItemlistgetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTxpItemItemlistgetRequest) GetApiMethodName() string {
    return "taobao.txp.item.itemlistget"
}

func (r TaobaoTxpItemItemlistgetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTxpItemItemlistgetRequest) SetBeginPage(beginPage int64) error {
    r.beginPage = beginPage
    r.Set("begin_page", beginPage)
    return nil
}

func (r TaobaoTxpItemItemlistgetRequest) GetBeginPage() int64 {
    return r.beginPage
}

func (r *TaobaoTxpItemItemlistgetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoTxpItemItemlistgetRequest) GetPageSize() int64 {
    return r.pageSize
}

