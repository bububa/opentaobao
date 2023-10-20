package lstbm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstBmStoreUpdateAPIResponse 修改品牌商自有门店数据 API返回值
// alibaba.lst.bm.store.update
//
// 修改品牌商自有门店数据
type AlibabaLstBmStoreUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaLstBmStoreUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstBmStoreUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstBmStoreUpdateAPIResponseModel).Reset()
}

// AlibabaLstBmStoreUpdateAPIResponseModel is 修改品牌商自有门店数据 成功返回结果
type AlibabaLstBmStoreUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_bm_store_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true表示执行成功，false表示执行失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstBmStoreUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaLstBmStoreUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstBmStoreUpdateAPIResponse)
	},
}

// GetAlibabaLstBmStoreUpdateAPIResponse 从 sync.Pool 获取 AlibabaLstBmStoreUpdateAPIResponse
func GetAlibabaLstBmStoreUpdateAPIResponse() *AlibabaLstBmStoreUpdateAPIResponse {
	return poolAlibabaLstBmStoreUpdateAPIResponse.Get().(*AlibabaLstBmStoreUpdateAPIResponse)
}

// ReleaseAlibabaLstBmStoreUpdateAPIResponse 将 AlibabaLstBmStoreUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstBmStoreUpdateAPIResponse(v *AlibabaLstBmStoreUpdateAPIResponse) {
	v.Reset()
	poolAlibabaLstBmStoreUpdateAPIResponse.Put(v)
}
