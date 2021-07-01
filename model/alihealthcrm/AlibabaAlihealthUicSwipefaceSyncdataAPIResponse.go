package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
刷脸测睡眠数据同步 API返回值 
alibaba.alihealth.uic.swipeface.syncdata

刷脸测睡眠数据同步，三方数据回传
*/
type AlibabaAlihealthUicSwipefaceSyncdataAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthUicSwipefaceSyncdataAPIResponseModel
}

// 刷脸测睡眠数据同步 成功返回结果
type AlibabaAlihealthUicSwipefaceSyncdataAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_uic_swipeface_syncdata_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // true调用成功，false失败
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
