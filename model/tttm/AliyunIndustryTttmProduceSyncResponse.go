package tttm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖生产进度同步 APIResponse
aliyun.industry.tttm.produce.sync

天天特卖生产进度同步
*/
type AliyunIndustryTttmProduceSyncAPIResponse struct {
    model.CommonResponse
    AliyunIndustryTttmProduceSyncResponse
}

type AliyunIndustryTttmProduceSyncResponse struct {
    XMLName xml.Name `xml:"aliyun_industry_tttm_produce_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 状态
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
