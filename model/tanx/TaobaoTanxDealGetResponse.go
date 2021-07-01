package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
对外部dsp提供交易id查询接口 API返回值 
taobao.tanx.deal.get

对外部dsp提供交易id查询接口
*/
type TaobaoTanxDealGetAPIResponse struct {
    model.CommonResponse
    TaobaoTanxDealGetResponse
}

// 对外部dsp提供交易id查询接口 成功返回结果
type TaobaoTanxDealGetResponse struct {
    XMLName xml.Name `xml:"tanx_deal_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果代码
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // 结果信息
    Messag   string `json:"messag,omitempty" xml:"messag,omitempty"`
    // 查询结果
    Sucess   bool `json:"sucess,omitempty" xml:"sucess,omitempty"`
    // 查询结果
    Result   *DealInfoDto `json:"result,omitempty" xml:"result,omitempty"`
}
