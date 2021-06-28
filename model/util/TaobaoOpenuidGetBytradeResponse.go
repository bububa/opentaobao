package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过订单获取对应买家的openUID APIResponse
taobao.openuid.get.bytrade

通过订单获取对应买家的openUID,需要卖家授权
*/
type TaobaoOpenuidGetBytradeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"openuid_get_bytrade_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 当前交易tid对应买家的openuid
    
    OpenUid   string `json:"open_uid,omitempty" xml:"