package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询底导数据 API返回值 
alibaba.alihealth.pregnancy.navigateinfo.query

备孕管理--获取底部导航信息
*/
type AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponseModel
}

// 查询底导数据 成功返回结果
type AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_navigateinfo_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果集
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
