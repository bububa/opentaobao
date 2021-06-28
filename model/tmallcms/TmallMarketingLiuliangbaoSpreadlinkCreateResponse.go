package tmallcms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建流量宝活动链接 APIResponse
tmall.marketing.liuliangbao.spreadlink.create

通过源活动链接和商家ID，创建流量宝活动链接
*/
type TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse struct {
    model.CommonResponse
    // Response *TmallMarketingLiuliangbaoSpreadlinkCreateResponse `json:"tmall_marketing_liuliangbao_spreadlink_create_response,omitempty"` 
    TmallMarketingLiuliangbaoSpreadlinkCreateResponse
}

/* model for simplify = false
type TmallMarketingLiuliangbaoSpreadlinkCreateResponse struct {

    // 流量宝系统执行结果
    
    Llbresult  *struct {
        LLBApiResult  *LLBApiResult `json:"llb_api_result,omitempty"`
    } `json:"llbresult,omitempty"`
    

}
*/

type TmallMarketingLiuliangbaoSpreadlinkCreateResponse struct {

    // 流量宝系统执行结果
    Llbresult   *LLBApiResult `json:"llbresult,omitempty"`

}
