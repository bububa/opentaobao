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

// New
