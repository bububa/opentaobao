package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
access token 获取精灵用户 id API返回值 
alibaba.ailab.user.open.uid.get

access token 获取精灵用户 id
*/
type AlibabaAilabUserOpenUidGetAPIResponse struct {
    model.CommonResponse
    AlibabaAilabUserOpenUidGetAPIResponseModel
}

// access token 获取精灵用户 id 成功返回结果
type AlibabaAilabUserOpenUidGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ailab_user_open_uid_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 详细信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // user id
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
    // 状态码，200 成功，其他失败
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}
