package lstbm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstBmStoreAddAPIResponse 导入品牌商自有门店 API返回值
// alibaba.lst.bm.store.add
//
// 导入品牌商自有门店
type AlibabaLstBmStoreAddAPIResponse struct {
	model.CommonResponse
	AlibabaLstBmStoreAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstBmStoreAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstBmStoreAddAPIResponseModel).Reset()
}

// AlibabaLstBmStoreAddAPIResponseModel is 导入品牌商自有门店 成功返回结果
type AlibabaLstBmStoreAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_bm_store_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true表示执行成功，false表示执行失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstBmStoreAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaLstBmStoreAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstBmStoreAddAPIResponse)
	},
}

// GetAlibabaLstBmStoreAddAPIResponse 从 sync.Pool 获取 AlibabaLstBmStoreAddAPIResponse
func GetAlibabaLstBmStoreAddAPIResponse() *AlibabaLstBmStoreAddAPIResponse {
	return poolAlibabaLstBmStoreAddAPIResponse.Get().(*AlibabaLstBmStoreAddAPIResponse)
}

// ReleaseAlibabaLstBmStoreAddAPIResponse 将 AlibabaLstBmStoreAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstBmStoreAddAPIResponse(v *AlibabaLstBmStoreAddAPIResponse) {
	v.Reset()
	poolAlibabaLstBmStoreAddAPIResponse.Put(v)
}
