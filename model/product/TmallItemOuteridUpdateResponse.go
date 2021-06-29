package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫商品/SKU商家编码更新接口 APIResponse
tmall.item.outerid.update

天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名）
*/
type TmallItemOuteridUpdateAPIResponse struct {
    model.CommonResponse
    TmallItemOuteridUpdateResponse
}

type TmallItemOuteridUpdateResponse struct {
    XMLName xml.Name `xml:"tmall_item_outerid_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 商家编码更新结果
    
    OuteridUpdateResult   string `json:"outerid_update_result,omitempty" xml:"outerid_update_result,omitempty"`

    
}
