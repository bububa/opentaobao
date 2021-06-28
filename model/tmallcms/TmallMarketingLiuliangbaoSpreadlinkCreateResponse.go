package tmallcms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建流量宝活动链接 APIResponse
tmall.marketing.liuliangbao.spreadlink.create

通过源活动链接和商家ID，创建流量宝活动链接
*/
type TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse struct {
    model.CommonResponse
    TmallMarketingLiuliangbaoSpreadlinkCreateResponse
}

type TmallMarketingLiuliangbaoSpreadlinkCreateResponse struct {
    XMLName xml.Name `xml:"tmall_marketing_liuliangbao_spreadlink_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 流量宝系统执行结果
    
    Llbresult   *LLBApiResult `json:"llbresult,omitempty" xml:"llbresult,omitempty"`

    
}
