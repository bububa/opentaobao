package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
艺人详情 API请求
alibaba.xiami.api.artist.detail.get

艺人详情
*/
type AlibabaXiamiApiArtistDetailGetRequest struct {
    model.Params
    // 艺人id
    _id   int64
    // 是否显示description, show为显示, 其他为不显示
    _description   string
}

// 初始化AlibabaXiamiApiArtistDetailGetRequest对象
func NewAlibabaXiamiApiArtistDetailGetRequest() *AlibabaXiamiApiArtistDetailGetRequest{
    return &AlibabaXiamiApiArtistDetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiArtistDetailGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.artist.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiArtistDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 艺人id
func (r *AlibabaXiamiApiArtistDetailGetRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaXiamiApiArtistDetailGetRequest) GetId() int64 {
    return r._id
}
// Description Setter
// 是否显示description, show为显示, 其他为不显示
func (r *AlibabaXiamiApiArtistDetailGetRequest) SetDescription(_description string) error {
    r._description = _description
    r.Set("description", _description)
    return nil
}

// Description Getter
func (r AlibabaXiamiApiArtistDetailGetRequest) GetDescription() string {
    return r._description
}
