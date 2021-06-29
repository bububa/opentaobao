package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼多媒体上传接口 API请求
alibaba.idle.rent.media.upload

上传多媒体信息，包括图片、视频（暂不支持）
*/
type AlibabaIdleRentMediaUploadRequest struct {
    model.Params
    // 多媒体字节数组
    _data   []*model.File
    // 文件名
    _name   string
    // 0-表示图片，1-表示视频（暂不支持）
    _type   int64
}

// 初始化AlibabaIdleRentMediaUploadRequest对象
func NewAlibabaIdleRentMediaUploadRequest() *AlibabaIdleRentMediaUploadRequest{
    return &AlibabaIdleRentMediaUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentMediaUploadRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.media.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentMediaUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Data Setter
// 多媒体字节数组
func (r *AlibabaIdleRentMediaUploadRequest) SetData(_data []*model.File) error {
    r._data = _data
    r.Set("data", _data)
    return nil
}

// Data Getter
func (r AlibabaIdleRentMediaUploadRequest) GetData() []*model.File {
    return r._data
}
// Name Setter
// 文件名
func (r *AlibabaIdleRentMediaUploadRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlibabaIdleRentMediaUploadRequest) GetName() string {
    return r._name
}
// Type Setter
// 0-表示图片，1-表示视频（暂不支持）
func (r *AlibabaIdleRentMediaUploadRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaIdleRentMediaUploadRequest) GetType() int64 {
    return r._type
}
