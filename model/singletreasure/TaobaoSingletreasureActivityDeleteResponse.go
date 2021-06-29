package singletreasure

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除活动接口 API返回值 
taobao.singletreasure.activity.delete

删除优惠活动
*/
type TaobaoSingletreasureActivityDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoSingletreasureActivityDeleteResponse
}

// 删除活动接口 成功返回结果
type TaobaoSingletreasureActivityDeleteResponse struct {
    XMLName xml.Name `xml:"singletreasure_activity_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 删除是否成功 boolean值
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`
    // 错误码
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 系统执行是否成功
    ResultStatus   bool `json:"result_status,omitempty" xml:"result_status,omitempty"`
}
