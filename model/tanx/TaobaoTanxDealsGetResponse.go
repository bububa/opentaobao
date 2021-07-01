package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取交易列表 API返回值 
taobao.tanx.deals.get

批量获取交易信息
*/
type TaobaoTanxDealsGetAPIResponse struct {
    model.CommonResponse
    TaobaoTanxDealsGetResponse
}

// 批量获取交易列表 成功返回结果
type TaobaoTanxDealsGetResponse struct {
    XMLName xml.Name `xml:"tanx_deals_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询是否成功
    Sucess   bool `json:"sucess,omitempty" xml:"sucess,omitempty"`
    // 查询结果编码
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // 查询结果信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 查询结果
    Deals   []DealInfoDto `json:"deals,omitempty" xml:"deals>deal_info_dto,omitempty"`
}
