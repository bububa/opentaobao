package tttm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖商品信息同步 APIResponse
aliyun.industry.tttm.items.sync

天天特卖商品信息同步
*/
type AliyunIndustryTttmItemsSyncAPIResponse struct {
    model.CommonResponse
    AliyunIndustryTttmItemsSyncResponse
}

type AliyunIndustryTttmItemsSyncResponse struct {
    XMLName xml.Name `xml:"aliyun_industry_tttm_items_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 成功失败标识
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
