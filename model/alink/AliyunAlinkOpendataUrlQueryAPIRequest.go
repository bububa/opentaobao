package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunAlinkOpendataUrlQueryAPIRequest
开放数据授权访问URL查询 API请求
aliyun.alink.opendata.url.query

厂商数据授权访问URL查询 */
type AliyunAlinkOpendataUrlQueryAPIRequest struct {
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

// NewAliyunAlinkOpendataUrlQueryRequest 初始化AliyunAlinkOpendataUrlQueryAPIRequest对象
func NewAliyunAlinkOpendataUrlQueryRequest() *AliyunAlinkOpendataUrlQueryAPIRequest {
	return &AliyunAlinkOpendataUrlQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunAlinkOpendataUrlQueryAPIRequest) GetApiMethodName() string {
	return "aliyun.alink.opendata.url.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunAlinkOpendataUrlQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AccessKey Setter
// 授权key，厂家在物联平台申请的云端授权key
func (r *AliyunAlinkOpendataUrlQueryAPIRequest) SetAccessKey(_accessKey string) error {
	r._accessKey = _accessKey
	r.Set("access_key", _accessKey)
	return nil
}

// Get AccessKey Getter
func (r AliyunAlinkOpendataUrlQueryAPIRequest) GetAccessKey() string {
	return r._accessKey
}

// Set is BizDay Setter
// 数据日期，格式：yyyyMMdd
func (r *AliyunAlinkOpendataUrlQueryAPIRequest) SetBizDay(_bizDay string) error {
	r._bizDay = _bizDay
	r.Set("biz_day", _bizDay)
	return nil
}

// Get BizDay Getter
func (r AliyunAlinkOpendataUrlQueryAPIRequest) GetBizDay() string {
	return r._bizDay
}

// Set is BizHour Setter
// 数据时点，范围[0,23]
func (r *AliyunAlinkOpendataUrlQueryAPIRequest) SetBizHour(_bizHour int64) error {
	r._bizHour = _bizHour
	r.Set("biz_hour", _bizHour)
	return nil
}

// Get BizHour Getter
func (r AliyunAlinkOpendataUrlQueryAPIRequest) GetBizHour() int64 {
	return r._bizHour
}

// Set is DataType Setter
// 数据类型，1：设备数据，2：用户操作数据
func (r *AliyunAlinkOpendataUrlQueryAPIRequest) SetDataType(_dataType int64) error {
	r._dataType = _dataType
	r.Set("data_type", _dataType)
	return nil
}

// Get DataType Getter
func (r AliyunAlinkOpendataUrlQueryAPIRequest) GetDataType() int64 {
	return r._dataType
}
