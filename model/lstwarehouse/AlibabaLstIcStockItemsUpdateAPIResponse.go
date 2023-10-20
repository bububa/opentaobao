package lstwarehouse

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaLstIcStockItemsUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstIcStockItemsUpdateAPIResponseModel).Reset()
}

// AlibabaLstIcStockItemsUpdateAPIResponseModel is 零售通经销商商品库存设置 成功返回结果
type AlibabaLstIcStockItemsUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_ic_stock_items_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLstIcStockItemsUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstIcStockItemsUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstIcStockItemsUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstIcStockItemsUpdateAPIResponse)
	},
}

// GetAlibabaLstIcStockItemsUpdateAPIResponse 从 sync.Pool 获取 AlibabaLstIcStockItemsUpdateAPIResponse
func GetAlibabaLstIcStockItemsUpdateAPIResponse() *AlibabaLstIcStockItemsUpdateAPIResponse {
	return poolAlibabaLstIcStockItemsUpdateAPIResponse.Get().(*AlibabaLstIcStockItemsUpdateAPIResponse)
}

// ReleaseAlibabaLstIcStockItemsUpdateAPIResponse 将 AlibabaLstIcStockItemsUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstIcStockItemsUpdateAPIResponse(v *AlibabaLstIcStockItemsUpdateAPIResponse) {
	v.Reset()
	poolAlibabaLstIcStockItemsUpdateAPIResponse.Put(v)
}
