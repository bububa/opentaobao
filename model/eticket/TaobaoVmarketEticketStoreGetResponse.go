package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取电子凭证预约门店信息 APIResponse
taobao.vmarket.eticket.store.get

用于给外部商家查询电子凭证预约门店信息
*/
type TaobaoVmarketEticketStoreGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"vmarket_eticket_store_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商户id
    
    StoreId   int64 `json:"store_id,omitempty" xml:"