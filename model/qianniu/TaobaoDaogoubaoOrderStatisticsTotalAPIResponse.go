package qianniu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售订单总额统计 API返回值 
taobao.daogoubao.order.statistics.total

对接千牛端数字中心
*/
type TaobaoDaogoubaoOrderStatisticsTotalAPIResponse struct {
    model.CommonResponse
    TaobaoDaogoubaoOrderStatisticsTotalAPIResponseModel
}

// 销售订单总额统计 成功返回结果
type TaobaoDaogoubaoOrderStatisticsTotalAPIResponseModel struct {
    XMLName xml.Name `xml:"daogoubao_order_statistics_total_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *OrderStatisticsResult `json:"result,omitempty" xml:"result,omitempty"`
}
