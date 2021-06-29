package store

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商户门店库创建接口 APIResponse
taobao.place.storegroup.create

用于商家创建线下门店库
*/
type TaobaoPlaceStoregroupCreateAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStoregroupCreateResponse
}

type TaobaoPlaceStoregroupCreateResponse struct {
    XMLName xml.Name `xml:"place_storegroup_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
