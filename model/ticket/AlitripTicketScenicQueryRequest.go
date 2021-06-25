package ticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】卖家已发布门票商品列表查询接口（根据景点维度查询） APIRequest
alitrip.ticket.scenic.query

查询卖家已发布过的门票商品列表，根据景点维度聚合查询。如果卖家在该景点下未发布过任何商品，则查询不到数据！
*/
type AlitripTicketScenicQueryRequest struct {
    model.Params

    // 标准景点ID。ali_scenic_id，out_scenic_id二者至少需要填写一个
    aliScenicId   int64 

    // 当前分页。每页默认最多返回20条数据
    currentPage   int64 

    // 商家景点ID。ali_scenic_id，out_scenic_id二者至少需要填写一个
    outScenicId   string 

}

func NewAlitripTicketScenicQueryRequest() *AlitripTicketScenicQueryRequest{
    return &AlitripTicketScenicQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTicketScenicQueryRequest) GetApiMethodName() string {
    return "alitrip.ticket.scenic.query"
}

func (r AlitripTicketScenicQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTicketScenicQueryRequest) SetAliScenicId(aliScenicId int64) error {
    r.aliScenicId = aliScenicId
    r.Set("ali_scenic_id", aliScenicId)
    return nil
}

func (r AlitripTicketScenicQueryRequest) GetAliScenicId() int64 {
    return r.aliScenicId
}

func (r *AlitripTicketScenicQueryRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlitripTicketScenicQueryRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlitripTicketScenicQueryRequest) SetOutScenicId(outScenicId string) error {
    r.outScenicId = outScenicId
    r.Set("out_scenic_id", outScenicId)
    return nil
}

func (r AlitripTicketScenicQueryRequest) GetOutScenicId() string {
    return r.outScenicId
}

