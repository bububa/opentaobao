package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增门店扩展属性 APIResponse
taobao.place.store.extend.add

新增授权用户的门店扩展属性
*/
type TaobaoPlaceStoreExtendAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"place_store_extend_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否失败
    
    Failure   bool `json:"failure,omitempty" xml:"