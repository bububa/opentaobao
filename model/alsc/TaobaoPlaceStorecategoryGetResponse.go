package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取门店类目信息 APIResponse
taobao.place.storecategory.get

获取门店类目信息
*/
type TaobaoPlaceStorecategoryGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"place_storecategory_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 门店类目格式
    
    CategoryList   string `json:"category_list,omitempty" xml:"