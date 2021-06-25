package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tanx竞价失败反馈api APIResponse
taobao.tanx.biddingrefuses.get

竞价失败反馈根据创意id查询API提供
*/
type TaobaoTanxBiddingrefusesGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTanxBiddingrefusesGetResponse `json:"taobao_tanx_biddingrefuses_get_response,omitempty"`
}

type TaobaoTanxBiddingrefusesGetResponse struct {

    // 返回是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 返回竞价失败对象信息
    BiddingRefuseList   []BiddingRefuseDto `json:"bidding_refuse_list,omitempty"`

}
