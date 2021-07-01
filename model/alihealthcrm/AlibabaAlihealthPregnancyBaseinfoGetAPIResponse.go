package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
拉取备孕初始化信息 API返回值 
alibaba.alihealth.pregnancy.baseinfo.get

拉取备孕初始化信息
*/
type AlibabaAlihealthPregnancyBaseinfoGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthPregnancyBaseinfoGetAPIResponseModel
}

// 拉取备孕初始化信息 成功返回结果
type AlibabaAlihealthPregnancyBaseinfoGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_baseinfo_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果集
    Result   *AlibabaAlihealthPregnancyBaseinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
