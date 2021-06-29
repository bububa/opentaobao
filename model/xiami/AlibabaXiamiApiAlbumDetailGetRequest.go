package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
虾米音乐专辑详情接口 API请求
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

// 初始化AlibabaXiamiApiAlbumDetailGetRequest对象
func NewAlibabaXiamiApiAlbumDetailGetRequest() *AlibabaXiamiApiAlbumDetailGetRequest{
    return &AlibabaXiamiApiAlbumDetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiAlbumDetailGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.album.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiAlbumDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 专辑ID
func (r *AlibabaXiamiApiAlbumDetailGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r AlibabaXiamiApiAlbumDetailGetRequest) GetId() int64 {
    return r.id
}
// FullDes Setter
// 是否获取完整描述
func (r *AlibabaXiamiApiAlbumDetailGetRequest) SetFullDes(fullDes bool) error {
    r.fullDes = fullDes
    r.Set("full_des", fullDes)
    return nil
}

// FullDes Getter
func (r AlibabaXiamiApiAlbumDetailGetRequest) GetFullDes() bool {
    return r.fullDes
}
