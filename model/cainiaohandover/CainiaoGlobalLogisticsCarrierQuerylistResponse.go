package cainiaohandover

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
实际承运商查询 APIResponse
cainiao.global.logistics.carrier.querylist

查询出所有的实际承运商
*/
type CainiaoGlobalLogisticsCarrierQuerylistAPIResponse struct {
    model.CommonResponse
    CainiaoGlobalLogisticsCarrierQuerylistResponse
}

type CainiaoGlobalLogisticsCarrierQuerylistResponse struct {
    XMLName xml.Name `xml:"cainiao_global_logistics_carrier_querylist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 1
    
    Result   *DubboResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
