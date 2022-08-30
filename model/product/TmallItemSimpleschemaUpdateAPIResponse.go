package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSimpleschemaUpdateAPIResponse 天猫简化编辑商品 API返回值
// tmall.item.simpleschema.update
//
// 国外大商家天猫简化编辑商品
type TmallItemSimpleschemaUpdateAPIResponse struct {
	model.CommonResponse
	TmallItemSimpleschemaUpdateAPIResponseModel
}

// TmallItemSimpleschemaUpdateAPIResponseModel is 天猫简化编辑商品 成功返回结果
type TmallItemSimpleschemaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_simpleschema_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 编辑商品的itemid
	UpdateItemResult string `json:"update_item_result,omitempty" xml:"update_item_result,omitempty"`
	// 编辑商品操作成功时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// sku与outerId映射信息
	SkuMapJson string `json:"sku_map_json,omitempty" xml:"sku_map_json,omitempty"`
}
