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
type AlibabaXiamiApiCollectDetailGetRequest struct {
    model.Params
    // 精选集ID
    id   int64
    // 是否获取完整描述
    fullDes   bool
    // 是否需要tag, show为需要, 其他为不需要
    tag   string
}

// 初始化AlibabaXiamiApiCollectDetailGetRequest对象
func NewAlibabaXiamiApiCollectDetailGetRequest() *AlibabaXiamiApiCollectDetailGetRequest{
    return &AlibabaXiamiApiCollectDetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiCollectDetailGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.collect.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiCollectDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 精选集ID
func (r *AlibabaXiamiApiCollectDetailGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r AlibabaXiamiApiCollectDetailGetRequest) GetId() int64 {
    return r.id
}
// FullDes Setter
// 是否获取完整描述
func (r *AlibabaXiamiApiCollectDetailGetRequest) SetFullDes(fullDes bool) error {
    r.fullDes = fullDes
    r.Set("full_des", fullDes)
    return nil
}

// FullDes Getter
func (r AlibabaXiamiApiCollectDetailGetRequest) GetFullDes() bool {
    return r.fullDes
}
// Tag Setter
// 是否需要tag, show为需要, 其他为不需要
func (r *AlibabaXiamiApiCollectDetailGetRequest) SetTag(tag string) error {
    r.tag = tag
    r.Set("tag", tag)
    return nil
}

// Tag Getter
func (r AlibabaXiamiApiCollectDetailGetRequest) GetTag() string {
    return r.tag
}
