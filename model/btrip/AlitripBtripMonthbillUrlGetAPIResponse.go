package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
月账单数据查询 API返回值 
alitrip.btrip.monthbill.url.get

月账单数据查询
*/
type AlitripBtripMonthbillUrlGetAPIResponse struct {
    model.CommonResponse
    AlitripBtripMonthbillUrlGetAPIResponseModel
}

// 月账单数据查询 成功返回结果
type AlitripBtripMonthbillUrlGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_btrip_monthbill_url_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果对象
    Results   []OpenAccountRs `json:"results,omitempty" xml:"results>open_account_rs,omitempty"`
    // 结果码
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 结果信息
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}
