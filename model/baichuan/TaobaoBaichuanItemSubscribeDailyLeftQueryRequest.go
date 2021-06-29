package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询当天可添加的余量 API请求
taobao.baichuan.item.subscribe.daily.left.query

查询当天可添加的余量
*/
type TaobaoBaichuanItemSubscribeDailyLeftQueryRequest struct {
    model.Params
}

// 初始化TaobaoBaichuanItemSubscribeDailyLeftQueryRequest对象
func NewTaobaoBaichuanItemSubscribeDailyLeftQueryRequest() *TaobaoBaichuanItemSubscribeDailyLeftQueryRequest{
    return &TaobaoBaichuanItemSubscribeDailyLeftQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemSubscribeDailyLeftQueryRequest) GetApiMethodName() string {
    return "taobao.baichuan.item.subscribe.daily.left.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemSubscribeDailyLeftQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
