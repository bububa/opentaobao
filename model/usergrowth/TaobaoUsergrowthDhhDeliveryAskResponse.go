package usergrowth

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
广告曝光前判定接口V2 APIResponse
taobao.usergrowth.dhh.delivery.ask

提供给媒体在曝光广告前调用
*/
type TaobaoUsergrowthDhhDeliveryAskAPIResponse struct {
    model.CommonResponse
    TaobaoUsergrowthDhhDeliveryAskResponse
}

type TaobaoUsergrowthDhhDeliveryAskResponse struct {
    XMLName xml.Name `xml:"usergrowth_dhh_delivery_ask_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // true: 目标用户；false: 非目标用户
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // 在大航海平台推荐的任务ID
    
    TaskId   string `json:"task_id,omitempty" xml:"task_id,omitempty"`

    
    // 错误码， 0： 成功；1：限流；2：服务不可用
    
    Errcode   int64 `json:"errcode,omitempty" xml:"errcode,omitempty"`

    
    // 在大航海平台可投放的任务ID列表
    
    TaskIdList   []string `json:"task_id_list,omitempty" xml:"task_id_list>string,omitempty"`
    
    
}
