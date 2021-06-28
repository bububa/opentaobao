package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据交易单号判断是否为菜鸟发货订单 APIResponse
cainiao.cboss.workplatform.logistics.iscainiaoorder

根据交易单号判断是否为菜鸟发货订单
*/
type CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_cboss_workplatform_logistics_iscainiaoorder_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // isCainiaoOrder
    
    IsCainiaoOrder   bool `json:"is_cainiao_order,omitempty" xml:"