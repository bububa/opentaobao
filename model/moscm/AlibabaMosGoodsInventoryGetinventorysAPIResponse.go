package moscm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsInventoryGetinventorysAPIResponse 可售库存查询 API返回值
// alibaba.mos.goods.inventory.getinventorys
//
// 查询商品的可售、在库和占库数量
type AlibabaMosGoodsInventoryGetinventorysAPIResponse struct {
	model.CommonResponse
	AlibabaMosGoodsInventoryGetinventorysAPIResponseModel
}

// AlibabaMosGoodsInventoryGetinventorysAPIResponseModel is 可售库存查询 成功返回结果
type AlibabaMosGoodsInventoryGetinventorysAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_goods_inventory_getinventorys_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据
	Datas []VirtualInventoryDto `json:"datas,omitempty" xml:"datas>virtual_inventory_dto,omitempty"`
}
