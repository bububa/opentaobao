package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票品接口pushItem APIRequest
alibaba.damai.mev.open.pushitem

开放接口 推送票品
*/
type AlibabaDamaiMevOpenPushitemRequest struct {
    model.Params

    // 入参pushItemParam
    pushItemParam   *PushTicketItemPushOpenParam 

}

func NewAlibabaDamaiMevOpenPushitemRequest() *AlibabaDamaiMevOpenPushitemRequest{
    return &AlibabaDamaiMevOpenPushitemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenPushitemRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushitem"
}

func (r AlibabaDamaiMevOpenPushitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenPushitemRequest) SetPushItemParam(pushItemParam *PushTicketItemPushOpenParam) error {
    r.pushItemParam = pushItemParam
    r.Set("push_item_param", pushItemParam)
    return nil
}

func (r AlibabaDamaiMevOpenPushitemRequest) GetPushItemParam() *PushTicketItemPushOpenParam {
    return r.pushItemParam
}

