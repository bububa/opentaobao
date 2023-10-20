package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvmediauploadAPIRequest 闲鱼服务商-图片上传 API请求
// alibaba.idle.isv.media.upload
//
// 供外部服务商ISV进行闲鱼商品发布时上传商品所需图片
type AlibabaidleisvmediauploadAPIRequest struct {
	model.Params
	// 文件名
	_name string
	// 多媒体字节数组
	_data *model.File
	// 0-表示图片，1-表示视频（暂不支持）
	_type int64
}

// NewAlibabaidleisvmediauploadRequest 初始化AlibabaidleisvmediauploadAPIRequest对象
func NewAlibabaidleisvmediauploadRequest() *AlibabaidleisvmediauploadAPIRequest {
	return &AlibabaidleisvmediauploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvmediauploadAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.media.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvmediauploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvmediauploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 文件名
func (r *AlibabaidleisvmediauploadAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaidleisvmediauploadAPIRequest) GetName() string {
	return r._name
}

// SetData is Data Setter
// 多媒体字节数组
func (r *AlibabaidleisvmediauploadAPIRequest) SetData(_data *model.File) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaidleisvmediauploadAPIRequest) GetData() *model.File {
	return r._data
}

// SetType is Type Setter
// 0-表示图片，1-表示视频（暂不支持）
func (r *AlibabaidleisvmediauploadAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaidleisvmediauploadAPIRequest) GetType() int64 {
	return r._type
}
