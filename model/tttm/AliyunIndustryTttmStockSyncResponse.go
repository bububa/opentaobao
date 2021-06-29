package tttm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖库存同步接口 APIResponse
aliyun.industry.tttm.stock.sync

天天特卖库存同步接口
*/
type AliyunIndustryTttmStockSyncAPIResponse struct {
    model.CommonResponse
    AliyunIndustryTttmStockSyncResponse
}

type AliyunIndustryTttmStockSyncResponse struct {
    XMLName xml.Name `xml:"aliyun_industry_tttm_stock_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 成功失败标识
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
