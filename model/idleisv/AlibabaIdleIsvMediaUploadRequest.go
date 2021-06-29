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
    _data   *model.File
    // 文件名
    _name   string
    // 0-表示图片，1-表示视频（暂不支持）
    _type   int64
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
func (r *AlibabaIdleIsvMediaUploadRequest) SetData(_data *model.File) error {
    r._data = _data
    r.Set("data", _data)
    return nil
}

// Data Getter
func (r AlibabaIdleIsvMediaUploadRequest) GetData() *model.File {
    return r._data
}
// Name Setter
// 文件名
func (r *AlibabaIdleIsvMediaUploadRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlibabaIdleIsvMediaUploadRequest) GetName() string {
    return r._name
}
// Type Setter
// 0-表示图片，1-表示视频（暂不支持）
func (r *AlibabaIdleIsvMediaUploadRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaIdleIsvMediaUploadRequest) GetType() int64 {
    return r._type
}
