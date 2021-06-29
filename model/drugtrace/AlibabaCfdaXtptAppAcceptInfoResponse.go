package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
协同平台数据下行接口 APIResponse
alibaba.cfda.xtpt.app.accept.info

协同平台数据下行接口
*/
type AlibabaCfdaXtptAppAcceptInfoAPIResponse struct {
    model.CommonResponse
    AlibabaCfdaXtptAppAcceptInfoResponse
}

type AlibabaCfdaXtptAppAcceptInfoResponse struct {
    XMLName xml.Name `xml:"alibaba_cfda_xtpt_app_accept_info_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaCfdaXtptAppAcceptInfoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
