package oversea

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汇率信息获取 APIResponse
alibaba.oversea.exchagerate.get

提供外部汇率查询接口
*/
type AlibabaOverseaExchagerateGetAPIResponse struct {
    model.CommonResponse
    AlibabaOverseaExchagerateGetResponse
}

type AlibabaOverseaExchagerateGetResponse struct {
    XMLName xml.Name `xml:"alibaba_oversea_exchagerate_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果描述
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
