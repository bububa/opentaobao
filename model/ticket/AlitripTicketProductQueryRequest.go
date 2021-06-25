package ticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】门票商品查询接口 APIRequest
alitrip.ticket.product.query

门票商品查询接口：返回商家上传的门票商品信息
*/
type AlitripTicketProductQueryRequest struct {
    model.Params

    // 商户自定义收费项目编码。与ali_product_id，item_id 三者至少填写一个
    outProductId   string 

    // 阿里标准收费项目id。与out_product_id，item_id 三者至少填写一个
    aliProductId   int64 

    // 商品ID。与out_product_id，ali_product_id三者至少填写一个
    itemId   int64 

    // 代表业务来源，gongxiao代表供销平台业务
    pageSource   string 

}

func NewAlitripTicketProductQueryRequest() *AlitripTicketProductQueryRequest{
    return &AlitripTicketProductQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTicketProductQueryRequest) GetApiMethodName() string {
    return "alitrip.ticket.product.query"
}

func (r AlitripTicketProductQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTicketProductQueryRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

func (r AlitripTicketProductQueryRequest) GetOutProductId() string {
    return r.outProductId
}

func (r *AlitripTicketProductQueryRequest) SetAliProductId(aliProductId int64) error {
    r.aliProductId = aliProductId
    r.Set("ali_product_id", aliProductId)
    return nil
}

func (r AlitripTicketProductQueryRequest) GetAliProductId() int64 {
    return r.aliProductId
}

func (r *AlitripTicketProductQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlitripTicketProductQueryRequest) GetItemId() int64 {
    return r.itemId
}

func (r *AlitripTicketProductQueryRequest) SetPageSource(pageSource string) error {
    r.pageSource = pageSource
    r.Set("page_source", pageSource)
    return nil
}

func (r AlitripTicketProductQueryRequest) GetPageSource() string {
    return r.pageSource
}

