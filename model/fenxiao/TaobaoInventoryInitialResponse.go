package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
库存初始化 APIResponse
taobao.inventory.initial

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账。
*/
type TaobaoInventoryInitialAPIResponse struct {
    model.CommonResponse
    TaobaoInventoryInitialResponse
}

type TaobaoInventoryInitialResponse struct {
    XMLName xml.Name `xml:"inventory_initial_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 提示信息
    
    TipInfos   []TipInfo `json:"tip_infos,omitempty" xml:"tip_infos>tip_info,omitempty"`
    
    
}
