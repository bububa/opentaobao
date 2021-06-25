package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询当天可添加的余量 APIRequest
taobao.baichuan.item.subscribe.daily.left.query

查询当天可添加的余量
*/
type TaobaoBaichuanItemSubscribeDailyLeftQueryRequest struct {
    model.Params

}

func NewTaobaoBaichuanItemSubscribeDailyLeftQueryRequest() *TaobaoBaichuanItemSubscribeDailyLeftQueryRequest{
    return &TaobaoBaichuanItemSubscribeDailyLeftQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanItemSubscribeDailyLeftQueryRequest) GetApiMethodName() string {
    return "taobao.baichuan.item.subscribe.daily.left.query"
}

func (r TaobaoBaichuanItemSubscribeDailyLeftQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


