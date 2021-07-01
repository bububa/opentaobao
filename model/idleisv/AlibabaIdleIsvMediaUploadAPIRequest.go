package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvMediaUploadAPIRequest
闲鱼服务商-图片上传 API请求
alibaba.idle.isv.media.upload

供外部服务商ISV进行闲鱼商品发布时上传商品所需图片 */
type AlibabaIdleIsvMediaUploadAPIRequest struct {
	model.Params
	// 多媒体字节数组
	_data *model.File
	// 文件名
	_name string
	// 0-表示图片，1-表示视频（暂不支持）
	_type int64
}

// NewAlibabaIdleIsvMediaUploadRequest 初始化AlibabaIdleIsvMediaUploadAPIRequest对象
func NewAlibabaIdleIsvMediaUploadRequest() *AlibabaIdleIsvMediaUploadAPIRequest {
	return &AlibabaIdleIsvMediaUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvMediaUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.media.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvMediaUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Data Setter
// 多媒体字节数组
func (r *AlibabaIdleIsvMediaUploadAPIRequest) SetData(_data *model.File) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// Get Data Getter
func (r AlibabaIdleIsvMediaUploadAPIRequest) GetData() *model.File {
	return r._data
}

// Set is Name Setter
// 文件名
func (r *AlibabaIdleIsvMediaUploadAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r AlibabaIdleIsvMediaUploadAPIRequest) GetName() string {
	return r._name
}

// Set is Type Setter
// 0-表示图片，1-表示视频（暂不支持）
func (r *AlibabaIdleIsvMediaUploadAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r AlibabaIdleIsvMediaUploadAPIRequest) GetType() int64 {
	return r._type
}
