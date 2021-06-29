package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
价格推送接口（批量全量） APIResponse
taobao.xhotel.rates.update

酒店产品库rate批量更新房态信息
*/
type TaobaoXhotelRatesUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRatesUpdateResponse
}

type TaobaoXhotelRatesUpdateResponse struct {
    XMLName xml.Name `xml:"xhotel_rates_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // gid_and_rateplan_ids
    
    GidAndRpids   []string `json:"gid_and_rpids,omitempty" xml:"gid_and_rpids>string,omitempty"`
    
    
}
