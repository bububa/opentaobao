package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkshopgetAPIResponse 淘宝客-推广者-店铺搜索 API返回值
// taobao.tbk.shop.get
//
// 淘宝客店铺查询
type TaobaotbkshopgetAPIResponse struct {
	model.CommonResponse
	TaobaotbkshopgetAPIResponseModel
}

// TaobaotbkshopgetAPIResponseModel is 淘宝客-推广者-店铺搜索 成功返回结果
type TaobaotbkshopgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_shop_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘宝客店铺
	Results []NtbkShop `json:"results,omitempty" xml:"results>ntbk_shop,omitempty"`
	// 搜索到符合条件的结果总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
