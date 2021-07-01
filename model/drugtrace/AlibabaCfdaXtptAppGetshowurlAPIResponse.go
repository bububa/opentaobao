package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
协同平台码查询页面url API返回值 
alibaba.cfda.xtpt.app.getshowurl

协同平台码查询页面url
*/
type AlibabaCfdaXtptAppGetshowurlAPIResponse struct {
    model.CommonResponse
    AlibabaCfdaXtptAppGetshowurlAPIResponseModel
}

// 协同平台码查询页面url 成功返回结果
type AlibabaCfdaXtptAppGetshowurlAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_cfda_xtpt_app_getshowurl_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否成功
    MsgSuccess   bool `json:"msg_success,omitempty" xml:"msg_success,omitempty"`
    // model
    Model   string `json:"model,omitempty" xml:"model,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 状态
    HttpStatusCode   int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
}
