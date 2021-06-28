package lifeservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店关系新增 APIResponse
taobao.place.store.relation.add

新增授权用户的门店关系信息
*/
type TaobaoPlaceStoreRelationAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"place_store_relation_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否失败
    
    Failure   bool `json:"failure,omitempty" xml:"