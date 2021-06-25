package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
虾米音乐专辑详情接口 APIRequest
alibaba.xiami.api.album.detail.get

虾米音乐专辑详情接口
*/
type AlibabaXiamiApiAlbumDetailGetRequest struct {
    model.Params

    // 专辑ID
    id   int64 

    // 是否获取完整描述
    fullDes   bool 

}

func NewAlibabaXiamiApiAlbumDetailGetRequest() *AlibabaXiamiApiAlbumDetailGetRequest{
    return &AlibabaXiamiApiAlbumDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiAlbumDetailGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.album.detail.get"
}

func (r AlibabaXiamiApiAlbumDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiAlbumDetailGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaXiamiApiAlbumDetailGetRequest) GetId() int64 {
    return r.id
}

func (r *AlibabaXiamiApiAlbumDetailGetRequest) SetFullDes(fullDes bool) error {
    r.fullDes = fullDes
    r.Set("full_des", fullDes)
    return nil
}

func (r AlibabaXiamiApiAlbumDetailGetRequest) GetFullDes() bool {
    return r.fullDes
}

