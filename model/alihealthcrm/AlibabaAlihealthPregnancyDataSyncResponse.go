package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
四类数据同步 APIResponse
alibaba.alihealth.pregnancy.data.sync

经期调整；基础体温；排卵试纸；B超测排数据同步
*/
type AlibabaAlihealthPregnancyDataSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthPregnancyDataSyncResponse
}

type AlibabaAlihealthPregnancyDataSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_data_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
}
