package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
薄荷同步三方记录 API返回值 
alibaba.fmhealth.weight.lossplan.syncweightdata

用于三方薄荷同步数据到健康会员
*/
type AlibabaFmhealthWeightLossplanSyncweightdataAPIResponse struct {
    model.CommonResponse
    AlibabaFmhealthWeightLossplanSyncweightdataResponse
}

// 薄荷同步三方记录 成功返回结果
type AlibabaFmhealthWeightLossplanSyncweightdataResponse struct {
    XMLName xml.Name `xml:"alibaba_fmhealth_weight_lossplan_syncweightdata_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
