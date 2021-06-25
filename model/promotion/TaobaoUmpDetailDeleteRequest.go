package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除活动详情 APIRequest
taobao.ump.detail.delete

删除活动详情
*/
type TaobaoUmpDetailDeleteRequest struct {
    model.Params

    // 活动详情id
    detailId   int64 

}

func NewTaobaoUmpDetailDeleteRequest() *TaobaoUmpDetailDeleteRequest{
    return &TaobaoUmpDetailDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpDetailDeleteRequest) GetApiMethodName() string {
    return "taobao.ump.detail.delete"
}

func (r TaobaoUmpDetailDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpDetailDeleteRequest) SetDetailId(detailId int64) error {
    r.detailId = detailId
    r.Set("detail_id", detailId)
    return nil
}

func (r TaobaoUmpDetailDeleteRequest) GetDetailId() int64 {
    return r.detailId
}

