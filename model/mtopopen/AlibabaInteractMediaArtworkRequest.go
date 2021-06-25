package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
原图相关鉴权接口 APIRequest
alibaba.interact.media.artwork

拍摄并上传原图相关鉴权接口
*/
type AlibabaInteractMediaArtworkRequest struct {
    model.Params

    // 系统自动生成
    id   string 

}

func NewAlibabaInteractMediaArtworkRequest() *AlibabaInteractMediaArtworkRequest{
    return &AlibabaInteractMediaArtworkRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractMediaArtworkRequest) GetApiMethodName() string {
    return "alibaba.interact.media.artwork"
}

func (r AlibabaInteractMediaArtworkRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractMediaArtworkRequest) SetId(id string) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaInteractMediaArtworkRequest) GetId() string {
    return r.id
}

