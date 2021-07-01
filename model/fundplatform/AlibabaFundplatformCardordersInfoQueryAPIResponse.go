package fundplatform

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据制卡单分页查询卡信息 API返回值 
alibaba.fundplatform.cardorders.info.query

该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
*/
type AlibabaFundplatformCardordersInfoQueryAPIResponse struct {
    model.CommonResponse
    AlibabaFundplatformCardordersInfoQueryAPIResponseModel
}

// 根据制卡单分页查询卡信息 成功返回结果
type AlibabaFundplatformCardordersInfoQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_fundplatform_cardorders_info_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *CardMakingInfoQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
