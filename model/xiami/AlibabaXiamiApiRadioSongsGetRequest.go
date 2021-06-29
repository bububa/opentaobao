package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取电台歌曲 API请求
alibaba.xiami.api.radio.songs.get

获取电台songs
*/
type AlibabaXiamiApiRadioSongsGetRequest struct {
    model.Params
    // 电台类型
    _type   int64
    // 电台id
    _oid   int64
    // 歌曲个数, 不传为20
    _limit   int64
}

// 初始化AlibabaXiamiApiRadioSongsGetRequest对象
func NewAlibabaXiamiApiRadioSongsGetRequest() *AlibabaXiamiApiRadioSongsGetRequest{
    return &AlibabaXiamiApiRadioSongsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiRadioSongsGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.radio.songs.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiRadioSongsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 电台类型
func (r *AlibabaXiamiApiRadioSongsGetRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaXiamiApiRadioSongsGetRequest) GetType() int64 {
    return r._type
}
// Oid Setter
// 电台id
func (r *AlibabaXiamiApiRadioSongsGetRequest) SetOid(_oid int64) error {
    r._oid = _oid
    r.Set("oid", _oid)
    return nil
}

// Oid Getter
func (r AlibabaXiamiApiRadioSongsGetRequest) GetOid() int64 {
    return r._oid
}
// Limit Setter
// 歌曲个数, 不传为20
func (r *AlibabaXiamiApiRadioSongsGetRequest) SetLimit(_limit int64) error {
    r._limit = _limit
    r.Set("limit", _limit)
    return nil
}

// Limit Getter
func (r AlibabaXiamiApiRadioSongsGetRequest) GetLimit() int64 {
    return r._limit
}
