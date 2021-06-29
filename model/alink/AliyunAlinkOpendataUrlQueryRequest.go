package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开放数据授权访问URL查询 API请求
aliyun.alink.opendata.url.query

厂商数据授权访问URL查询
*/
type AliyunAlinkOpendataUrlQueryRequest struct {
    model.Params
    // 授权key，厂家在物联平台申请的云端授权key
    accessKey   string
    // 数据日期，格式：yyyyMMdd
    bizDay   string
    // 数据时点，范围[0,23]
    bizHour   int64
    // 数据类型，1：设备数据，2：用户操作数据
    dataType   int64
}

// 初始化AliyunAlinkOpendataUrlQueryRequest对象
func NewAliyunAlinkOpendataUrlQueryRequest() *AliyunAlinkOpendataUrlQueryRequest{
    return &AliyunAlinkOpendataUrlQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunAlinkOpendataUrlQueryRequest) GetApiMethodName() string {
    return "aliyun.alink.opendata.url.query"
}

// IRequest interface 方法, 获取API参数
func (r AliyunAlinkOpendataUrlQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccessKey Setter
// 授权key，厂家在物联平台申请的云端授权key
func (r *AliyunAlinkOpendataUrlQueryRequest) SetAccessKey(accessKey string) error {
    r.accessKey = accessKey
    r.Set("access_key", accessKey)
    return nil
}

// AccessKey Getter
func (r AliyunAlinkOpendataUrlQueryRequest) GetAccessKey() string {
    return r.accessKey
}
// BizDay Setter
// 数据日期，格式：yyyyMMdd
func (r *AliyunAlinkOpendataUrlQueryRequest) SetBizDay(bizDay string) error {
    r.bizDay = bizDay
    r.Set("biz_day", bizDay)
    return nil
}

// BizDay Getter
func (r AliyunAlinkOpendataUrlQueryRequest) GetBizDay() string {
    return r.bizDay
}
// BizHour Setter
// 数据时点，范围[0,23]
func (r *AliyunAlinkOpendataUrlQueryRequest) SetBizHour(bizHour int64) error {
    r.bizHour = bizHour
    r.Set("biz_hour", bizHour)
    return nil
}

// BizHour Getter
func (r AliyunAlinkOpendataUrlQueryRequest) GetBizHour() int64 {
    return r.bizHour
}
// DataType Setter
// 数据类型，1：设备数据，2：用户操作数据
func (r *AliyunAlinkOpendataUrlQueryRequest) SetDataType(dataType int64) error {
    r.dataType = dataType
    r.Set("data_type", dataType)
    return nil
}

// DataType Getter
func (r AliyunAlinkOpendataUrlQueryRequest) GetDataType() int64 {
    return r.dataType
}
