package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动详情列表 APIRequest
taobao.ump.details.get

分页查询优惠详情列表
*/
type TaobaoUmpDetailsGetRequest struct {
    model.Params

    // 营销活动id
    actId   int64 

    // 分页的页码
    pageNo   int64 

    // 每页的最大条数
    pageSize   int64 

}

func NewTaobaoUmpDetailsGetRequest() *TaobaoUmpDetailsGetRequest{
    return &TaobaoUmpDetailsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpDetailsGetRequest) GetApiMethodName() string {
    return "taobao.ump.details.get"
}

func (r TaobaoUmpDetailsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpDetailsGetRequest) SetActId(actId int64) error {
    r.actId = actId
    r.Set("act_id", actId)
    return nil
}

func (r TaobaoUmpDetailsGetRequest) GetActId() int64 {
    return r.actId
}

func (r *TaobaoUmpDetailsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoUmpDetailsGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoUmpDetailsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoUmpDetailsGetRequest) GetPageSize() int64 {
    return r.pageSize
}

