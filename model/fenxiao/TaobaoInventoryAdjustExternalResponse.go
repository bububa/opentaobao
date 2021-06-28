package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
非交易库存调整单 APIResponse
taobao.inventory.adjust.external

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家非交易调整库存，调拨出库、盘点等时调用
*/
type TaobaoInventoryAdjustExternalAPIResponse struct {
    model.CommonResponse
    TaobaoInventoryAdjustExternalResponse
}

type TaobaoInventoryAdjustExternalResponse struct {
    XMLName xml.Name `xml:"inventory_adjust_external_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作返回码
    
    OperateCode   string `json:"operate_code,omitempty" xml:"operate_code,omitempty"`

    
    // 提示信息
    
    TipInfos   []TipInfo `json:"tip_infos,omitempty" xml:"tip_infos>tip_info,omitempty"`
    
    
}
