package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleReportMediaUploadAPIRequest 验货报告上传文件 API请求
// alibaba.idle.report.media.upload
//
// 服务商上传文件、图片
type AlibabaIdleReportMediaUploadAPIRequest struct {
	model.Params
	// 名称
	_name string
	// 类型 PHOTO 图片；FILM 视频
	_type string
	// 多媒体字节数组
	_data *model.File
}

// NewAlibabaIdleReportMediaUploadRequest 初始化AlibabaIdleReportMediaUploadAPIRequest对象
func NewAlibabaIdleReportMediaUploadRequest() *AlibabaIdleReportMediaUploadAPIRequest {
	return &AlibabaIdleReportMediaUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleReportMediaUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.report.media.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleReportMediaUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleReportMediaUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 名称
func (r *AlibabaIdleReportMediaUploadAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaIdleReportMediaUploadAPIRequest) GetName() string {
	return r._name
}

// SetType is Type Setter
// 类型 PHOTO 图片；FILM 视频
func (r *AlibabaIdleReportMediaUploadAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaIdleReportMediaUploadAPIRequest) GetType() string {
	return r._type
}

// SetData is Data Setter
// 多媒体字节数组
func (r *AlibabaIdleReportMediaUploadAPIRequest) SetData(_data *model.File) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaIdleReportMediaUploadAPIRequest) GetData() *model.File {
	return r._data
}
