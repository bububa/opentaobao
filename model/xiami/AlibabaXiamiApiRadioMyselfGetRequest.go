package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
我的电台 API请求
alibaba.xiami.api.radio.myself.get

我的电台
*/
type AlibabaXiamiApiRadioMyselfGetRequest struct {
    model.Params
    // 歌曲数量
    _limit   int64
}

// 初始化AlibabaXiamiApiRadioMyselfGetRequest对象
func NewAlibabaXiamiApiRadioMyselfGetRequest() *AlibabaXiamiApiRadioMyselfGetRequest{
    return &AlibabaXiamiApiRadioMyselfGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiRadioMyselfGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.radio.myself.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiRadioMyselfGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Limit Setter
// 歌曲数量
func (r *AlibabaXiamiApiRadioMyselfGetRequest) SetLimit(_limit int64) error {
    r._limit = _limit
    r.Set("limit", _limit)
    return nil
}

// Limit Getter
func (r AlibabaXiamiApiRadioMyselfGetRequest) GetLimit() int64 {
    return r._limit
}
