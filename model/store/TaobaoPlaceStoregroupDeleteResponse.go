package store

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除门店库 APIResponse
taobao.place.storegroup.delete

删除门店库
*/
type TaobaoPlaceStoregroupDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStoregroupDeleteResponse
}

type TaobaoPlaceStoregroupDeleteResponse struct {
    XMLName xml.Name `xml:"place_storegroup_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
