package usergrowth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘系用增线下转化明细 APIResponse
taobao.usergrowth.offline.convertion.details.get

淘系用增增长-线下拉新：为渠道提供返回拉新转化数据接口
*/
type TaobaoUsergrowthOfflineConvertionDetailsGetAPIResponse struct {
    model.CommonResponse
    TaobaoUsergrowthOfflineConvertionDetailsGetResponse
}

type TaobaoUsergrowthOfflineConvertionDetailsGetResponse struct {
    XMLName xml.Name `xml:"usergrowth_offline_convertion_details_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 鹰眼id
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`

    
    // 返回数据的总条数
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`

    
    // 每页多少条记录
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`

    
    // 页码
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`

    
    // 返回数据集合
    
    List   []TaobaoUsergrowthOfflineConvertionDetailsGetE `json:"list,omitempty" xml:"list>taobao_usergrowth_offline_convertion_details_get_e,omitempty"`
    
    
}
