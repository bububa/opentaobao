package wlbimports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
一般进口发货 APIResponse
taobao.wlb.imports.general.consign

将订单信息发送到菜鸟海外转运仓；
业务规则：
1）交易订单为待发货状态。
2）单笔订单多个商品，交易金额不能大于1000人民币。
*/
type TaobaoWlbImportsGeneralConsignAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_imports_general_consign_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 发货成功后的物流订单编号
    
    LgorderCode   string `json:"lgorder_code,omitempty" xml:"