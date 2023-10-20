package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentMediaUploadAPIRequest 闲鱼多媒体上传接口 API请求
// alibaba.idle.rent.media.upload
//
// 上传多媒体信息，包括图片、视频（暂不支持）
type AlibabaIdleRentMediaUploadAPIRequest struct {
	model.Params
	// 文件名
	_name string
	// 多媒体字节数组
	_data *model.File
	// 0-表示图片，1-表示视频（暂不支持）
	_type int64
}

// NewAlibabaIdleRentMediaUploadRequest 初始化AlibabaIdleRentMediaUploadAPIRequest对象
func NewAlibabaIdleRentMediaUploadRequest() *AlibabaIdleRentMediaUploadAPIRequest {
	return &AlibabaIdleRentMediaUploadAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRentMediaUploadAPIRequest) Reset() {
	r._name = ""
	r._data = nil
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentMediaUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.media.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentMediaUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRentMediaUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 文件名
func (r *AlibabaIdleRentMediaUploadAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaIdleRentMediaUploadAPIRequest) GetName() string {
	return r._name
}

// SetData is Data Setter
// 多媒体字节数组
func (r *AlibabaIdleRentMediaUploadAPIRequest) SetData(_data *model.File) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaIdleRentMediaUploadAPIRequest) GetData() *model.File {
	return r._data
}

// SetType is Type Setter
// 0-表示图片，1-表示视频（暂不支持）
func (r *AlibabaIdleRentMediaUploadAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaIdleRentMediaUploadAPIRequest) GetType() int64 {
	return r._type
}

var poolAlibabaIdleRentMediaUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRentMediaUploadRequest()
	},
}

// GetAlibabaIdleRentMediaUploadRequest 从 sync.Pool 获取 AlibabaIdleRentMediaUploadAPIRequest
func GetAlibabaIdleRentMediaUploadAPIRequest() *AlibabaIdleRentMediaUploadAPIRequest {
	return poolAlibabaIdleRentMediaUploadAPIRequest.Get().(*AlibabaIdleRentMediaUploadAPIRequest)
}

// ReleaseAlibabaIdleRentMediaUploadAPIRequest 将 AlibabaIdleRentMediaUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRentMediaUploadAPIRequest(v *AlibabaIdleRentMediaUploadAPIRequest) {
	v.Reset()
	poolAlibabaIdleRentMediaUploadAPIRequest.Put(v)
}
