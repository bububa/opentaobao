package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼服务商-图片上传 API请求
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

// 初始化AlibabaIdleIsvMediaUploadRequest对象
func NewAlibabaIdleIsvMediaUploadRequest() *AlibabaIdleIsvMediaUploadRequest{
    return &AlibabaIdleIsvMediaUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvMediaUploadRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.media.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvMediaUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Data Setter
// 多媒体字节数组
func (r *AlibabaIdleIsvMediaUploadRequest) SetData(data []*model.File) error {
    r.data = data
    r.Set("data", data)
    return nil
}

// Data Getter
func (r AlibabaIdleIsvMediaUploadRequest) GetData() []*model.File {
    return r.data
}
// Name Setter
// 文件名
func (r *AlibabaIdleIsvMediaUploadRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaIdleIsvMediaUploadRequest) GetName() string {
    return r.name
}
// Type Setter
// 0-表示图片，1-表示视频（暂不支持）
func (r *AlibabaIdleIsvMediaUploadRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaIdleIsvMediaUploadRequest) GetType() int64 {
    return r.type
}
