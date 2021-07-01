package fundplatform

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过卡号查询卡信息 API返回值 
alibaba.fundplatform.cardorders.info.query.by.cardno

该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
*/
type AlibabaFundplatformCardordersInfoQueryByCardnoAPIResponse struct {
    model.CommonResponse
    AlibabaFundplatformCardordersInfoQueryByCardnoAPIResponseModel
}

// 通过卡号查询卡信息 成功返回结果
type AlibabaFundplatformCardordersInfoQueryByCardnoAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_fundplatform_cardorders_info_query_by_cardno_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *CardMakingInfoQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
