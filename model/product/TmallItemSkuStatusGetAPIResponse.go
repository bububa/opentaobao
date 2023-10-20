package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallitemskustatusgetAPIResponse 商品sku上下架查询 API返回值
// tmall.item.sku.status.get
//
// 商品sku上下架状态查询
type TmallitemskustatusgetAPIResponse struct {
	model.CommonResponse
	TmallitemskustatusgetAPIResponseModel
}

// TmallitemskustatusgetAPIResponseModel is 商品sku上下架查询 成功返回结果
type TmallitemskustatusgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sku_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TmallitemskustatusgetApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
