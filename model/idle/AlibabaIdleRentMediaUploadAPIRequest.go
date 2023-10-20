package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentmediauploadAPIRequest 闲鱼多媒体上传接口 API请求
// alibaba.idle.rent.media.upload
//
// 上传多媒体信息，包括图片、视频（暂不支持）
type AlibabaidlerentmediauploadAPIRequest struct {
	model.Params
	// 文件名
	_name string
	// 多媒体字节数组
	_data *model.File
	// 0-表示图片，1-表示视频（暂不支持）
	_type int64
}

// NewAlibabaidlerentmediauploadRequest 初始化AlibabaidlerentmediauploadAPIRequest对象
func NewAlibabaidlerentmediauploadRequest() *AlibabaidlerentmediauploadAPIRequest {
	return &AlibabaidlerentmediauploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerentmediauploadAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.media.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerentmediauploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerentmediauploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 文件名
func (r *AlibabaidlerentmediauploadAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaidlerentmediauploadAPIRequest) GetName() string {
	return r._name
}

// SetData is Data Setter
// 多媒体字节数组
func (r *AlibabaidlerentmediauploadAPIRequest) SetData(_data *model.File) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaidlerentmediauploadAPIRequest) GetData() *model.File {
	return r._data
}

// SetType is Type Setter
// 0-表示图片，1-表示视频（暂不支持）
func (r *AlibabaidlerentmediauploadAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaidlerentmediauploadAPIRequest) GetType() int64 {
	return r._type
}
