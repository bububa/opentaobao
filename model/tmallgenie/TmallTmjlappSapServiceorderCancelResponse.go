package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消售后服务单 APIResponse
tmall.tmjlapp.sap.serviceorder.cancel

SAP跟天猫精灵app接口对接，用户在app取消sap售后服务工单
*/
type TmallTmjlappSapServiceorderCancelAPIResponse struct {
    model.CommonResponse
    TmallTmjlappSapServiceorderCancelResponse
}

type TmallTmjlappSapServiceorderCancelResponse struct {
    XMLName xml.Name `xml:"tmall_tmjlapp_sap_serviceorder_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 取消服务单返回
    
    CancelResponse   *Dtcancelresponse `json:"cancel_response,omitempty" xml:"cancel_response,omitempty"`

    
}
