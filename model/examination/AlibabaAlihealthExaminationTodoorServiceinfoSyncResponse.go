package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上门检测服务信息同步 APIResponse
alibaba.alihealth.examination.todoor.serviceinfo.sync

isv同步上门检测服务信息给健康
*/
type AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationTodoorServiceinfoSyncResponse
}

type AlibabaAlihealthExaminationTodoorServiceinfoSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_todoor_serviceinfo_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
