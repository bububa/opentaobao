package dt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAdsDataImportAPIResponse 数据导入 API返回值
// taobao.ads.data.import
//
// 数据导入
type TaobaoAdsDataImportAPIResponse struct {
	model.CommonResponse
	TaobaoAdsDataImportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAdsDataImportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAdsDataImportAPIResponseModel).Reset()
}

// TaobaoAdsDataImportAPIResponseModel is 数据导入 成功返回结果
type TaobaoAdsDataImportAPIResponseModel struct {
	XMLName xml.Name `xml:"ads_data_import_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 0:成功/-1:失败
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAdsDataImportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetMsg = ""
	m.RetCode = 0
}

var poolTaobaoAdsDataImportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAdsDataImportAPIResponse)
	},
}

// GetTaobaoAdsDataImportAPIResponse 从 sync.Pool 获取 TaobaoAdsDataImportAPIResponse
func GetTaobaoAdsDataImportAPIResponse() *TaobaoAdsDataImportAPIResponse {
	return poolTaobaoAdsDataImportAPIResponse.Get().(*TaobaoAdsDataImportAPIResponse)
}

// ReleaseTaobaoAdsDataImportAPIResponse 将 TaobaoAdsDataImportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAdsDataImportAPIResponse(v *TaobaoAdsDataImportAPIResponse) {
	v.Reset()
	poolTaobaoAdsDataImportAPIResponse.Put(v)
}
