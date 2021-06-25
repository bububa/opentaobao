package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
虾米音乐精选集详情接口 APIRequest
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

func NewAlibabaXiamiApiCollectDetailGetRequest() *AlibabaXiamiApiCollectDetailGetRequest{
    return &AlibabaXiamiApiCollectDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiCollectDetailGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.collect.detail.get"
}

func (r AlibabaXiamiApiCollectDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiCollectDetailGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaXiamiApiCollectDetailGetRequest) GetId() int64 {
    return r.id
}

func (r *AlibabaXiamiApiCollectDetailGetRequest) SetFullDes(fullDes bool) error {
    r.fullDes = fullDes
    r.Set("full_des", fullDes)
    return nil
}

func (r AlibabaXiamiApiCollectDetailGetRequest) GetFullDes() bool {
    return r.fullDes
}

func (r *AlibabaXiamiApiCollectDetailGetRequest) SetTag(tag string) error {
    r.tag = tag
    r.Set("tag", tag)
    return nil
}

func (r AlibabaXiamiApiCollectDetailGetRequest) GetTag() string {
    return r.tag
}

