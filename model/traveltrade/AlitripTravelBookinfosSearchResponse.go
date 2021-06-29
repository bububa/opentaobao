package traveltrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单预定信息列表搜索接口 APIResponse
alitrip.travel.bookinfos.search

查询订单预定信息列表
*/
type AlitripTravelBookinfosSearchAPIResponse struct {
    model.CommonResponse
    AlitripTravelBookinfosSearchResponse
}

type AlitripTravelBookinfosSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_bookinfos_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
    // 订单及预约id映射列表
    
    OrderBookInfoList   []FirstResult `json:"order_book_info_list,omitempty" xml:"order_book_info_list>first_result,omitempty"`
    
    
    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
    // true或false
    
    IsSuccess   string `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
