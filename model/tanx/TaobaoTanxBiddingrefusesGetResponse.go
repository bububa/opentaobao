package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tanx竞价失败反馈api APIResponse
taobao.tanx.biddingrefuses.get

竞价失败反馈根据创意id查询API提供
*/
type TaobaoTanxBiddingrefusesGetAPIResponse struct {
    model.CommonResponse
    TaobaoTanxBiddingrefusesGetResponse
}

type TaobaoTanxBiddingrefusesGetResponse struct {
    XMLName xml.Name `xml:"tanx_biddingrefuses_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 返回竞价失败对象信息
    
    BiddingRefuseList   []BiddingRefuseDto `json:"bidding_refuse_list,omitempty" xml:"bidding_refuse_list>bidding_refuse_dto,omitempty"`
    
    
}
