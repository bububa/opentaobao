package tuanhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitriptuanhotelitemskudeleteAPIResponse 酒店团购套餐商品SKU删除 API返回值
// alitrip.tuan.hotel.item.sku.delete
//
// 商户对发布的宝贝套餐价格库存信息进行删除
type AlitriptuanhotelitemskudeleteAPIResponse struct {
	model.CommonResponse
	AlitriptuanhotelitemskudeleteAPIResponseModel
}

// AlitriptuanhotelitemskudeleteAPIResponseModel is 酒店团购套餐商品SKU删除 成功返回结果
type AlitriptuanhotelitemskudeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tuan_hotel_item_sku_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// sku列表
	TopItemSkuBaseInfoList []TopItemSkuBaseInfoList `json:"top_item_sku_base_info_list,omitempty" xml:"top_item_sku_base_info_list>top_item_sku_base_info_list,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 宝贝ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 操作状态
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}
