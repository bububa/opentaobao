package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼多媒体上传接口 APIRequest
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

func NewAlibabaIdleRentMediaUploadRequest() *AlibabaIdleRentMediaUploadRequest{
    return &AlibabaIdleRentMediaUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRentMediaUploadRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.media.upload"
}

func (r AlibabaIdleRentMediaUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRentMediaUploadRequest) SetData(data []*model.File) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r AlibabaIdleRentMediaUploadRequest) GetData() []*model.File {
    return r.data
}

func (r *AlibabaIdleRentMediaUploadRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlibabaIdleRentMediaUploadRequest) GetName() string {
    return r.name
}

func (r *AlibabaIdleRentMediaUploadRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaIdleRentMediaUploadRequest) GetType() int64 {
    return r.type
}

