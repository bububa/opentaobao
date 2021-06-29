package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询底导数据 APIResponse
alibaba.alihealth.pregnancy.navigateinfo.query

备孕管理--获取底部导航信息
*/
type AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthPregnancyNavigateinfoQueryResponse
}

type AlibabaAlihealthPregnancyNavigateinfoQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_navigateinfo_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果集
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
