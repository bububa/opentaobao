package lstvending

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendingGoodsSaveAPIResponse 自动售卖机商品回传 API返回值
// alibaba.lst.vending.goods.save
//
// 零售通自动售卖机商品数据回流。
type AlibabaLstVendingGoodsSaveAPIResponse struct {
	model.CommonResponse
	AlibabaLstVendingGoodsSaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstVendingGoodsSaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstVendingGoodsSaveAPIResponseModel).Reset()
}

// AlibabaLstVendingGoodsSaveAPIResponseModel is 自动售卖机商品回传 成功返回结果
type AlibabaLstVendingGoodsSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_vending_goods_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *MultiResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstVendingGoodsSaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstVendingGoodsSaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendingGoodsSaveAPIResponse)
	},
}

// GetAlibabaLstVendingGoodsSaveAPIResponse 从 sync.Pool 获取 AlibabaLstVendingGoodsSaveAPIResponse
func GetAlibabaLstVendingGoodsSaveAPIResponse() *AlibabaLstVendingGoodsSaveAPIResponse {
	return poolAlibabaLstVendingGoodsSaveAPIResponse.Get().(*AlibabaLstVendingGoodsSaveAPIResponse)
}

// ReleaseAlibabaLstVendingGoodsSaveAPIResponse 将 AlibabaLstVendingGoodsSaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstVendingGoodsSaveAPIResponse(v *AlibabaLstVendingGoodsSaveAPIResponse) {
	v.Reset()
	poolAlibabaLstVendingGoodsSaveAPIResponse.Put(v)
}
