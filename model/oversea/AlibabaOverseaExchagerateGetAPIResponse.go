package oversea

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汇率信息获取 API返回值 
alibaba.oversea.exchagerate.get

提供外部汇率查询接口
*/
type AlibabaOverseaExchagerateGetAPIResponse struct {
    model.CommonResponse
    AlibabaOverseaExchagerateGetAPIResponseModel
}

// 汇率信息获取 成功返回结果
type AlibabaOverseaExchagerateGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_oversea_exchagerate_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询结果描述
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
