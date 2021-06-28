package cntms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟配商家仓库发货 APIResponse
cainiao.cntms.logistics.order.consign

商家包装打印面单结束后，通知菜鸟包裹要发货
*/
type CainiaoCntmsLogisticsOrderConsignAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_cntms_logistics_order_consign_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 物流单号
    
    LogisticsOrderCode   string `json:"logistics_order_code,omitempty" xml:"