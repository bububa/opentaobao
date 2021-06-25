package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动详情 APIRequest
taobao.ump.detail.get

查询活动详情
*/
type TaobaoUmpDetailGetRequest struct {
    model.Params

    // 活动详情的id
    detailId   int64 

}

func NewTaobaoUmpDetailGetRequest() *TaobaoUmpDetailGetRequest{
    return &TaobaoUmpDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpDetailGetRequest) GetApiMethodName() string {
    return "taobao.ump.detail.get"
}

func (r TaobaoUmpDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpDetailGetRequest) SetDetailId(detailId int64) error {
    r.detailId = detailId
    r.Set("detail_id", detailId)
    return nil
}

func (r TaobaoUmpDetailGetRequest) GetDetailId() int64 {
    return r.detailId
}

