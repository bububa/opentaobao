package lstwarehouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstIcStockItemsUpdateAPIResponse 零售通经销商商品库存设置 API返回值
// alibaba.lst.ic.stock.items.update
//
// 零售通经销商商品库存设置
type AlibabaLstIcStockItemsUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaLstIcStockItemsUpdateAPIResponseModel
}

// AlibabaLstIcStockItemsUpdateAPIResponseModel is 零售通经销商商品库存设置 成功返回结果
type AlibabaLstIcStockItemsUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_ic_stock_items_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLstIcStockItemsUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
