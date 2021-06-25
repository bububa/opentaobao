package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
艺人详情 APIRequest
alibaba.xiami.api.artist.detail.get

艺人详情
*/
type AlibabaXiamiApiArtistDetailGetRequest struct {
    model.Params

    // 艺人id
    id   int64 

    // 是否显示description, show为显示, 其他为不显示
    description   string 

}

func NewAlibabaXiamiApiArtistDetailGetRequest() *AlibabaXiamiApiArtistDetailGetRequest{
    return &AlibabaXiamiApiArtistDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiArtistDetailGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.artist.detail.get"
}

func (r AlibabaXiamiApiArtistDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiArtistDetailGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaXiamiApiArtistDetailGetRequest) GetId() int64 {
    return r.id
}

func (r *AlibabaXiamiApiArtistDetailGetRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

func (r AlibabaXiamiApiArtistDetailGetRequest) GetDescription() string {
    return r.description
}

