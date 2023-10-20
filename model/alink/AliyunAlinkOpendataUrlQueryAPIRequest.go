package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunalinkopendataurlqueryAPIRequest 开放数据授权访问URL查询 API请求
// aliyun.alink.opendata.url.query
//
// 厂商数据授权访问URL查询
type AliyunalinkopendataurlqueryAPIRequest struct {
	model.Params
	// 授权key，厂家在物联平台申请的云端授权key
	_accessKey string
	// 数据日期，格式：yyyyMMdd
	_bizDay string
	// 数据时点，范围[0,23]
	_bizHour int64
	// 数据类型，1：设备数据，2：用户操作数据
	_dataType int64
}

// NewAliyunalinkopendataurlqueryRequest 初始化AliyunalinkopendataurlqueryAPIRequest对象
func NewAliyunalinkopendataurlqueryRequest() *AliyunalinkopendataurlqueryAPIRequest {
	return &AliyunalinkopendataurlqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunalinkopendataurlqueryAPIRequest) GetApiMethodName() string {
	return "aliyun.alink.opendata.url.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunalinkopendataurlqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunalinkopendataurlqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessKey is AccessKey Setter
// 授权key，厂家在物联平台申请的云端授权key
func (r *AliyunalinkopendataurlqueryAPIRequest) SetAccessKey(_accessKey string) error {
	r._accessKey = _accessKey
	r.Set("access_key", _accessKey)
	return nil
}

// GetAccessKey AccessKey Getter
func (r AliyunalinkopendataurlqueryAPIRequest) GetAccessKey() string {
	return r._accessKey
}

// SetBizDay is BizDay Setter
// 数据日期，格式：yyyyMMdd
func (r *AliyunalinkopendataurlqueryAPIRequest) SetBizDay(_bizDay string) error {
	r._bizDay = _bizDay
	r.Set("biz_day", _bizDay)
	return nil
}

// GetBizDay BizDay Getter
func (r AliyunalinkopendataurlqueryAPIRequest) GetBizDay() string {
	return r._bizDay
}

// SetBizHour is BizHour Setter
// 数据时点，范围[0,23]
func (r *AliyunalinkopendataurlqueryAPIRequest) SetBizHour(_bizHour int64) error {
	r._bizHour = _bizHour
	r.Set("biz_hour", _bizHour)
	return nil
}

// GetBizHour BizHour Getter
func (r AliyunalinkopendataurlqueryAPIRequest) GetBizHour() int64 {
	return r._bizHour
}

// SetDataType is DataType Setter
// 数据类型，1：设备数据，2：用户操作数据
func (r *AliyunalinkopendataurlqueryAPIRequest) SetDataType(_dataType int64) error {
	r._dataType = _dataType
	r.Set("data_type", _dataType)
	return nil
}

// GetDataType DataType Getter
func (r AliyunalinkopendataurlqueryAPIRequest) GetDataType() int64 {
	return r._dataType
}
