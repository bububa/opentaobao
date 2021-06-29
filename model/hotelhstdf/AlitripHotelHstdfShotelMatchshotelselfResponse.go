package hotelhstdf

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自主匹配标准酒店以及卖家酒店 APIResponse
alitrip.hotel.hstdf.shotel.matchshotelself

商家通过指定的标准酒店id和卖家酒店id进行匹配
*/
type AlitripHotelHstdfShotelMatchshotelselfAPIResponse struct {
    model.CommonResponse
    AlitripHotelHstdfShotelMatchshotelselfResponse
}

type AlitripHotelHstdfShotelMatchshotelselfResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_matchshotelself_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码
    
    Errorcode   string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`

    
    // 错误信息
    
    Errormsg   string `json:"errormsg,omitempty" xml:"errormsg,omitempty"`

    
    // 是否成功
    
    Status   bool `json:"status,omitempty" xml:"status,omitempty"`

    
}
