package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品hscode信息审核状态查询接口 APIRequest
tmall.item.hscode.audit.results.query

通过此接口查询天猫跨境商品的hscode信息审核状态，卖家可以参考返回结果判断是否需要调整商品hscode相关信息。
*/
type TmallItemHscodeAuditResultsQueryRequest struct {
    model.Params

    // 商品ID
    itemId   int64 

}

func NewTmallItemHscodeAuditResultsQueryRequest() *TmallItemHscodeAuditResultsQueryRequest{
    return &TmallItemHscodeAuditResultsQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemHscodeAuditResultsQueryRequest) GetApiMethodName() string {
    return "tmall.item.hscode.audit.results.query"
}

func (r TmallItemHscodeAuditResultsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemHscodeAuditResultsQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallItemHscodeAuditResultsQueryRequest) GetItemId() int64 {
    return r.itemId
}

