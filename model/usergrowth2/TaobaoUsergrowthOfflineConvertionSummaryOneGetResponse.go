package usergrowth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘用增 线下业务 转化t+1汇总数据 APIResponse
taobao.usergrowth.offline.convertion.summary.one.get

淘系用户增长团队-线下拉新业务，对线下渠道提供mapi，目的是为了给有开发能力的渠道提供返数功能，方便渠道对手下的推广者结算（t+1汇总）
*/
type TaobaoUsergrowthOfflineConvertionSummaryOneGetAPIResponse struct {
    model.CommonResponse
    TaobaoUsergrowthOfflineConvertionSummaryOneGetResponse
}

type TaobaoUsergrowthOfflineConvertionSummaryOneGetResponse struct {
    XMLName xml.Name `xml:"usergrowth_offline_convertion_summary_one_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 跟踪id
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`

    
    // 总条数
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`

    
    // 每页多少条记录
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`

    
    // 1
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`

    
    // 集合
    
    List   []OfflineConvertionSummaryT1Dto `json:"list,omitempty" xml:"list>offline_convertion_summary_t1dto,omitempty"`
    
    
}
