package tttm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖数字工厂订单获取 API返回值 
aliyun.industry.tttm.order.query

获取阿里云数字工厂内天天特卖业务的订单
*/
type AliyunIndustryTttmOrderQueryAPIResponse struct {
    model.CommonResponse
    AliyunIndustryTttmOrderQueryResponse
}

// 天天特卖数字工厂订单获取 成功返回结果
type AliyunIndustryTttmOrderQueryResponse struct {
    XMLName xml.Name `xml:"aliyun_industry_tttm_order_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 订单
    Result   *OrderDTO `json:"result,omitempty" xml:"result,omitempty"`
}
