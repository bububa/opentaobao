package shenjing

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取OSS上传参数 API返回值 
alibaba.ib.shenjing.visitor.pad.getinfo

PAD 端获取OSS上传参数，向OSS服务器上传图片。
*/
type AlibabaIbShenjingVisitorPadGetinfoAPIResponse struct {
    model.CommonResponse
    AlibabaIbShenjingVisitorPadGetinfoAPIResponseModel
}

// 获取OSS上传参数 成功返回结果
type AlibabaIbShenjingVisitorPadGetinfoAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ib_shenjing_visitor_pad_getinfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 响应对象
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 返回状态码
    Content   *PostObjectPolicyVo `json:"content,omitempty" xml:"content,omitempty"`
    // 放回中文提示
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // request_id
    ResultRequestId   string `json:"result_request_id,omitempty" xml:"result_request_id,omitempty"`
    // 是否成功
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
