package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票品接口pushItem API请求
alibaba.damai.mev.open.pushitem

开放接口 推送票品
*/
type AlibabaDamaiMevOpenPushitemRequest struct {
    model.Params
    // 入参pushItemParam
    _pushItemParam   *PushTicketItemPushOpenParam
}

// 初始化AlibabaDamaiMevOpenPushitemRequest对象
func NewAlibabaDamaiMevOpenPushitemRequest() *AlibabaDamaiMevOpenPushitemRequest{
    return &AlibabaDamaiMevOpenPushitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushitemRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushItemParam Setter
// 入参pushItemParam
func (r *AlibabaDamaiMevOpenPushitemRequest) SetPushItemParam(_pushItemParam *PushTicketItemPushOpenParam) error {
    r._pushItemParam = _pushItemParam
    r.Set("push_item_param", _pushItemParam)
    return nil
}

// PushItemParam Getter
func (r AlibabaDamaiMevOpenPushitemRequest) GetPushItemParam() *PushTicketItemPushOpenParam {
    return r._pushItemParam
}
