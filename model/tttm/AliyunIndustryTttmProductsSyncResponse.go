package tttm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖货品信息同步 APIResponse
aliyun.industry.tttm.products.sync

天天特卖货品信息同步
*/
type AliyunIndustryTttmProductsSyncAPIResponse struct {
    model.CommonResponse
    AliyunIndustryTttmProductsSyncResponse
}

type AliyunIndustryTttmProductsSyncResponse struct {
    XMLName xml.Name `xml:"aliyun_industry_tttm_products_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 成功失败标识
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
