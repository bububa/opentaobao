package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
对外部dsp提供交易id查询接口 APIResponse
taobao.tanx.deal.get

对外部dsp提供交易id查询接口
*/
type TaobaoTanxDealGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTanxDealGetResponse `json:"tanx_deal_get_response,omitempty"` 
    TaobaoTanxDealGetResponse
}

/* model for simplify = false
type TaobaoTanxDealGetResponse struct {

    // 结果代码
    
    Code   int64 `json:"code,omitempty"`
    

    // 结果信息
    
    Messag   string `json:"messag,omitempty"`
    

    // 查询结果
    
    Sucess   bool `json:"sucess,omitempty"`
    

    // 查询结果
    
    Result  *struct {
        DealInfoDTO  *DealInfoDTO `json:"deal_info_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTanxDealGetResponse struct {

    // 结果代码
    Code   int64 `json:"code,omitempty"`

    // 结果信息
    Messag   string `json:"messag,omitempty"`

    // 查询结果
    Sucess   bool `json:"sucess,omitempty"`

    // 查询结果
    Result   *DealInfoDTO `json:"result,omitempty"`

}
