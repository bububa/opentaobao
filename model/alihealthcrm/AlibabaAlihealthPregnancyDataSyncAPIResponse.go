package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
四类数据同步 API返回值 
alibaba.alihealth.pregnancy.data.sync

经期调整；基础体温；排卵试纸；B超测排数据同步
*/
type AlibabaAlihealthPregnancyDataSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthPregnancyDataSyncAPIResponseModel
}

// 四类数据同步 成功返回结果
type AlibabaAlihealthPregnancyDataSyncAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_data_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
