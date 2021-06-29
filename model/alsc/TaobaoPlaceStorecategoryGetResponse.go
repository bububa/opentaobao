package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取门店类目信息 API返回值 
taobao.place.storecategory.get

获取门店类目信息
*/
type TaobaoPlaceStorecategoryGetAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStorecategoryGetResponse
}

// 获取门店类目信息 成功返回结果
type TaobaoPlaceStorecategoryGetResponse struct {
    XMLName xml.Name `xml:"place_storecategory_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 门店类目格式
    CategoryList   string `json:"category_list,omitempty" xml:"category_list,omitempty"`
}
