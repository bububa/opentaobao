package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
原图相关鉴权接口 API请求
alibaba.interact.media.artwork

拍摄并上传原图相关鉴权接口
*/
type AlibabaInteractMediaArtworkRequest struct {
    model.Params
    // 系统自动生成
    _id   string
}

// 初始化AlibabaInteractMediaArtworkRequest对象
func NewAlibabaInteractMediaArtworkRequest() *AlibabaInteractMediaArtworkRequest{
    return &AlibabaInteractMediaArtworkRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractMediaArtworkRequest) GetApiMethodName() string {
    return "alibaba.interact.media.artwork"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractMediaArtworkRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 系统自动生成
func (r *AlibabaInteractMediaArtworkRequest) SetId(_id string) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaInteractMediaArtworkRequest) GetId() string {
    return r._id
}
