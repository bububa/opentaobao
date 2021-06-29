package traveltrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单二次预约查询接口 APIResponse
alitrip.travel.bookinfo.query

飞猪度假订单二次预约详情查询接口
*/
type AlitripTravelBookinfoQueryAPIResponse struct {
    model.CommonResponse
    AlitripTravelBookinfoQueryResponse
}

type AlitripTravelBookinfoQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_bookinfo_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 交易预定结果对象
    
    FirstResult   *TopTripBookInfoResult `json:"first_result,omitempty" xml:"first_result,omitempty"`

    
}
