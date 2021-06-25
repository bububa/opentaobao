package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
我的电台 APIRequest
alibaba.xiami.api.radio.myself.get

我的电台
*/
type AlibabaXiamiApiRadioMyselfGetRequest struct {
    model.Params

    // 歌曲数量
    limit   int64 

}

func NewAlibabaXiamiApiRadioMyselfGetRequest() *AlibabaXiamiApiRadioMyselfGetRequest{
    return &AlibabaXiamiApiRadioMyselfGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiRadioMyselfGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.radio.myself.get"
}

func (r AlibabaXiamiApiRadioMyselfGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiRadioMyselfGetRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

func (r AlibabaXiamiApiRadioMyselfGetRequest) GetLimit() int64 {
    return r.limit
}

