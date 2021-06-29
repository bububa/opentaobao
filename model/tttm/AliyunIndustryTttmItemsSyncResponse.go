package tttm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖商品信息同步 API返回值 
aliyun.industry.tttm.items.sync

天天特卖商品信息同步
*/
type AliyunIndustryTttmItemsSyncAPIResponse struct {
    model.CommonResponse
    AliyunIndustryTttmItemsSyncResponse
}

// 天天特卖商品信息同步 成功返回结果
type AliyunIndustryTttmItemsSyncResponse struct {
    XMLName xml.Name `xml:"aliyun_industry_tttm_items_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 成功失败标识
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
