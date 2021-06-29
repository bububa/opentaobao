package moscm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
可售库存查询 APIResponse
alibaba.mos.goods.inventory.getinventorys

查询商品的可售、在库和占库数量
*/
type AlibabaMosGoodsInventoryGetinventorysAPIResponse struct {
    model.CommonResponse
    AlibabaMosGoodsInventoryGetinventorysResponse
}

type AlibabaMosGoodsInventoryGetinventorysResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_goods_inventory_getinventorys_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回的数据
    
    Datas   []VirtualInventoryDto `json:"datas,omitempty" xml:"datas>virtual_inventory_dto,omitempty"`
    
    
}
