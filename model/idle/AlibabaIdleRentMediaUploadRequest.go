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
    data   []*model.File
    // 文件名
    name   string
    // 0-表示图片，1-表示视频（暂不支持）
    type   int64
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
func (r *AlibabaIdleRentMediaUploadRequest) SetData(data []*model.File) error {
    r.data = data
    r.Set("data", data)
    return nil
}

// Data Getter
func (r AlibabaIdleRentMediaUploadRequest) GetData() []*model.File {
    return r.data
}
// Name Setter
// 文件名
func (r *AlibabaIdleRentMediaUploadRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaIdleRentMediaUploadRequest) GetName() string {
    return r.name
}
// Type Setter
// 0-表示图片，1-表示视频（暂不支持）
func (r *AlibabaIdleRentMediaUploadRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaIdleRentMediaUploadRequest) GetType() int64 {
    return r.type
}
