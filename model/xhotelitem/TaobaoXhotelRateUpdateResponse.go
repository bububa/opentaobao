package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
价格推送接口（全量更新） API返回值 
taobao.xhotel.rate.update

酒店产品库rate更新
*/
type TaobaoXhotelRateUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRateUpdateResponse
}

// 价格推送接口（全量更新） 成功返回结果
type TaobaoXhotelRateUpdateResponse struct {
    XMLName xml.Name `xml:"xhotel_rate_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 酒店商品ID-酒店RPid
    GidAndRpid   string `json:"gid_and_rpid,omitempty" xml:"gid_and_rpid,omitempty"`
}
