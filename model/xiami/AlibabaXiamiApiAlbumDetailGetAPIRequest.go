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
type AlibabaXiamiApiAlbumDetailGetAPIRequest struct {
    model.Params
    // 专辑ID
    _id   int64
    // 是否获取完整描述
    _fullDes   bool
}

// 初始化AlibabaXiamiApiAlbumDetailGetAPIRequest对象
func NewAlibabaXiamiApiAlbumDetailGetRequest() *AlibabaXiamiApiAlbumDetailGetAPIRequest{
    return &AlibabaXiamiApiAlbumDetailGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiAlbumDetailGetAPIRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.album.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiAlbumDetailGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 专辑ID
func (r *AlibabaXiamiApiAlbumDetailGetAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaXiamiApiAlbumDetailGetAPIRequest) GetId() int64 {
    return r._id
}
// FullDes Setter
// 是否获取完整描述
func (r *AlibabaXiamiApiAlbumDetailGetAPIRequest) SetFullDes(_fullDes bool) error {
    r._fullDes = _fullDes
    r.Set("full_des", _fullDes)
    return nil
}

// FullDes Getter
func (r AlibabaXiamiApiAlbumDetailGetAPIRequest) GetFullDes() bool {
    return r._fullDes
}
