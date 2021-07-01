package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDrugShopListAPIResponse
查询卖家外卖店列表 API返回值
taobao.drug.shop.list

查询卖家外卖店列表 */
type TaobaoDrugShopListAPIResponse struct {
	model.CommonResponse
	TaobaoDrugShopListAPIResponseModel
}

// TaobaoDrugShopListAPIResponseModel is 查询卖家外卖店列表 成功返回结果
type TaobaoDrugShopListAPIResponseModel struct {
	XMLName xml.Name `xml:"drug_shop_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据结果集
	Result *TakeoutShopPage `json:"result,omitempty" xml:"result,omitempty"`
}
