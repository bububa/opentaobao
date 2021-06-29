package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼服务商-图片上传 APIRequest
alibaba.idle.isv.media.upload

供外部服务商ISV进行闲鱼商品发布时上传商品所需图片
*/
type AlibabaIdleIsvMediaUploadRequest struct {
    model.Params

    // 多媒体字节数组
    data   []*model.File 

    // 文件名
    name   string 

    // 0-表示图片，1-表示视频（暂不支持）
    type   int64 

}

func NewAlibabaIdleIsvMediaUploadRequest() *AlibabaIdleIsvMediaUploadRequest{
    return &AlibabaIdleIsvMediaUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleIsvMediaUploadRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.media.upload"
}

func (r AlibabaIdleIsvMediaUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleIsvMediaUploadRequest) SetData(data []*model.File) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r AlibabaIdleIsvMediaUploadRequest) GetData() []*model.File {
    return r.data
}

func (r *AlibabaIdleIsvMediaUploadRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlibabaIdleIsvMediaUploadRequest) GetName() string {
    return r.name
}

func (r *AlibabaIdleIsvMediaUploadRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaIdleIsvMediaUploadRequest) GetType() int64 {
    return r.type
}

