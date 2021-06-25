package cntms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟配商家仓库发货 APIResponse
cainiao.cntms.logistics.order.consign

商家包装打印面单结束后，通知菜鸟包裹要发货
*/
type CainiaoCntmsLogisticsOrderConsignAPIResponse struct {
    model.CommonResponse
    Response *CainiaoCntmsLogisticsOrderConsignResponse `json:"cainiao_cntms_logistics_order_consign_response,omitempty"`
}

type CainiaoCntmsLogisticsOrderConsignResponse struct {

    // 物流单号
    LogisticsOrderCode   string `json:"logistics_order_code,omitempty"`

}
