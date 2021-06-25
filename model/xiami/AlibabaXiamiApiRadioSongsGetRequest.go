package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取电台歌曲 APIRequest
alibaba.xiami.api.radio.songs.get

获取电台songs
*/
type AlibabaXiamiApiRadioSongsGetRequest struct {
    model.Params

    // 电台类型
    type   int64 

    // 电台id
    oid   int64 

    // 歌曲个数, 不传为20
    limit   int64 

}

func NewAlibabaXiamiApiRadioSongsGetRequest() *AlibabaXiamiApiRadioSongsGetRequest{
    return &AlibabaXiamiApiRadioSongsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiRadioSongsGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.radio.songs.get"
}

func (r AlibabaXiamiApiRadioSongsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiRadioSongsGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaXiamiApiRadioSongsGetRequest) GetType() int64 {
    return r.type
}

func (r *AlibabaXiamiApiRadioSongsGetRequest) SetOid(oid int64) error {
    r.oid = oid
    r.Set("oid", oid)
    return nil
}

func (r AlibabaXiamiApiRadioSongsGetRequest) GetOid() int64 {
    return r.oid
}

func (r *AlibabaXiamiApiRadioSongsGetRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

func (r AlibabaXiamiApiRadioSongsGetRequest) GetLimit() int64 {
    return r.limit
}

