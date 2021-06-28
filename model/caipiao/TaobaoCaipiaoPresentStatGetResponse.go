package caipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家按天统计的彩票赠送数据 APIResponse
taobao.caipiao.present.stat.get

查询卖家一段时间内按天统计的彩票赠送数据，只支持查询90天以内的数据.
*/
type TaobaoCaipiaoPresentStatGetAPIResponse struct {
    model.CommonResponse
    TaobaoCaipiaoPresentStatGetResponse
}

type TaobaoCaipiaoPresentStatGetResponse struct {
    XMLName xml.Name `xml:"caipiao_present_stat_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询的结果集大小
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
    // 查询的结果集
    
    Results   []LotteryWangcaiPresentStat `json:"results,omitempty" xml:"results>lottery_wangcai_present_stat,omitempty"`
    
    
}
