package lstlogistics2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeSellerOfflineOrderUploadAPIResponse 供应商-线下订单-导入接口 API返回值
// alibaba.lst.trade.seller.offline.order.upload
//
// 供应商线下订单数据上传、实现和零售通本地云仓订单的共配
type AlibabaLstTradeSellerOfflineOrderUploadAPIResponse struct {
	model.CommonResponse
	AlibabaLstTradeSellerOfflineOrderUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstTradeSellerOfflineOrderUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstTradeSellerOfflineOrderUploadAPIResponseModel).Reset()
}

// AlibabaLstTradeSellerOfflineOrderUploadAPIResponseModel is 供应商-线下订单-导入接口 成功返回结果
type AlibabaLstTradeSellerOfflineOrderUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_seller_offline_order_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLstTradeSellerOfflineOrderUploadResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstTradeSellerOfflineOrderUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstTradeSellerOfflineOrderUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeSellerOfflineOrderUploadAPIResponse)
	},
}

// GetAlibabaLstTradeSellerOfflineOrderUploadAPIResponse 从 sync.Pool 获取 AlibabaLstTradeSellerOfflineOrderUploadAPIResponse
func GetAlibabaLstTradeSellerOfflineOrderUploadAPIResponse() *AlibabaLstTradeSellerOfflineOrderUploadAPIResponse {
	return poolAlibabaLstTradeSellerOfflineOrderUploadAPIResponse.Get().(*AlibabaLstTradeSellerOfflineOrderUploadAPIResponse)
}

// ReleaseAlibabaLstTradeSellerOfflineOrderUploadAPIResponse 将 AlibabaLstTradeSellerOfflineOrderUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstTradeSellerOfflineOrderUploadAPIResponse(v *AlibabaLstTradeSellerOfflineOrderUploadAPIResponse) {
	v.Reset()
	poolAlibabaLstTradeSellerOfflineOrderUploadAPIResponse.Put(v)
}
