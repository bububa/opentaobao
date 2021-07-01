package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询往来单位列表 API返回值 
alibaba.alihealth.drug.kyt.listparts

查询往来单位列表
*/
type AlibabaAlihealthDrugKytListpartsAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytListpartsAPIResponseModel
}

// 查询往来单位列表 成功返回结果
type AlibabaAlihealthDrugKytListpartsAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_listparts_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 监控宝推送网站监控信息，返回结果
    Result   *AlibabaAlihealthDrugKytListpartsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
