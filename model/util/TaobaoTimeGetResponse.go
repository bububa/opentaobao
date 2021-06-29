package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取淘宝系统当前时间 API返回值 
taobao.time.get

获取淘宝系统当前时间
*/
type TaobaoTimeGetAPIResponse struct {
    model.CommonResponse
    TaobaoTimeGetResponse
}

// 获取淘宝系统当前时间 成功返回结果
type TaobaoTimeGetResponse struct {
    XMLName xml.Name `xml:"time_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 淘宝系统当前时间。格式:yyyy-MM-dd HH:mm:ss
    Time   string `json:"time,omitempty" xml:"time,omitempty"`
}
