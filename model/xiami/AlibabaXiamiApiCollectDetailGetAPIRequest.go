package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
虾米音乐精选集详情接口 API请求
alibaba.xiami.api.collect.detail.get

虾米音乐精选集详情接口
*/
type AlibabaXiamiApiCollectDetailGetAPIRequest struct {
    model.Params
    // 精选集ID
    _id   int64
    // 是否获取完整描述
    _fullDes   bool
    // 是否需要tag, show为需要, 其他为不需要
    _tag   string
}

// 初始化AlibabaXiamiApiCollectDetailGetAPIRequest对象
func NewAlibabaXiamiApiCollectDetailGetRequest() *AlibabaXiamiApiCollectDetailGetAPIRequest{
    return &AlibabaXiamiApiCollectDetailGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiCollectDetailGetAPIRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.collect.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiCollectDetailGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 精选集ID
func (r *AlibabaXiamiApiCollectDetailGetAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaXiamiApiCollectDetailGetAPIRequest) GetId() int64 {
    return r._id
}
// FullDes Setter
// 是否获取完整描述
func (r *AlibabaXiamiApiCollectDetailGetAPIRequest) SetFullDes(_fullDes bool) error {
    r._fullDes = _fullDes
    r.Set("full_des", _fullDes)
    return nil
}

// FullDes Getter
func (r AlibabaXiamiApiCollectDetailGetAPIRequest) GetFullDes() bool {
    return r._fullDes
}
// Tag Setter
// 是否需要tag, show为需要, 其他为不需要
func (r *AlibabaXiamiApiCollectDetailGetAPIRequest) SetTag(_tag string) error {
    r._tag = _tag
    r.Set("tag", _tag)
    return nil
}

// Tag Getter
func (r AlibabaXiamiApiCollectDetailGetAPIRequest) GetTag() string {
    return r._tag
}
