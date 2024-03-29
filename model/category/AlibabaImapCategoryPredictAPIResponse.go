package category

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaImapCategoryPredictAPIResponse 类目预测接口 API返回值
// alibaba.imap.category.predict
//
// * 类目预测接口
//   - 【必填字段】 title, srcChannelId, srcCategoryId, targetChannelId
//   - 【非必填，但有最好填上】itemId, barcode, brandName, pvPairDOList, srcCatNamePathList
type AlibabaImapCategoryPredictAPIResponse struct {
	model.CommonResponse
	AlibabaImapCategoryPredictAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaImapCategoryPredictAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaImapCategoryPredictAPIResponseModel).Reset()
}

// AlibabaImapCategoryPredictAPIResponseModel is 类目预测接口 成功返回结果
type AlibabaImapCategoryPredictAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_imap_category_predict_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaImapCategoryPredictResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaImapCategoryPredictAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaImapCategoryPredictAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaImapCategoryPredictAPIResponse)
	},
}

// GetAlibabaImapCategoryPredictAPIResponse 从 sync.Pool 获取 AlibabaImapCategoryPredictAPIResponse
func GetAlibabaImapCategoryPredictAPIResponse() *AlibabaImapCategoryPredictAPIResponse {
	return poolAlibabaImapCategoryPredictAPIResponse.Get().(*AlibabaImapCategoryPredictAPIResponse)
}

// ReleaseAlibabaImapCategoryPredictAPIResponse 将 AlibabaImapCategoryPredictAPIResponse 保存到 sync.Pool
func ReleaseAlibabaImapCategoryPredictAPIResponse(v *AlibabaImapCategoryPredictAPIResponse) {
	v.Reset()
	poolAlibabaImapCategoryPredictAPIResponse.Put(v)
}
