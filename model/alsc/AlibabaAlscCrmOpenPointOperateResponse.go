package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
积分操作接口 APIResponse
alibaba.alsc.crm.open.point.operate

同步积分接口
*/
type AlibabaAlscCrmOpenPointOperateAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmOpenPointOperateResponse
}

type AlibabaAlscCrmOpenPointOperateResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_open_point_operate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
