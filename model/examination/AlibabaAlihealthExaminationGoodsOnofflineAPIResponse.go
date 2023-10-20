package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationGoodsOnofflineAPIResponse 上线/下线 体检产品 API返回值
// alibaba.alihealth.examination.goods.onoffline
//
// 第三方体检机构对接钉钉体检中的产品 上线／下线
type AlibabaAlihealthExaminationGoodsOnofflineAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationGoodsOnofflineAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationGoodsOnofflineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationGoodsOnofflineAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationGoodsOnofflineAPIResponseModel is 上线/下线 体检产品 成功返回结果
type AlibabaAlihealthExaminationGoodsOnofflineAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_goods_onoffline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationGoodsOnofflineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthExaminationGoodsOnofflineAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationGoodsOnofflineAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationGoodsOnofflineAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationGoodsOnofflineAPIResponse
func GetAlibabaAlihealthExaminationGoodsOnofflineAPIResponse() *AlibabaAlihealthExaminationGoodsOnofflineAPIResponse {
	return poolAlibabaAlihealthExaminationGoodsOnofflineAPIResponse.Get().(*AlibabaAlihealthExaminationGoodsOnofflineAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationGoodsOnofflineAPIResponse 将 AlibabaAlihealthExaminationGoodsOnofflineAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationGoodsOnofflineAPIResponse(v *AlibabaAlihealthExaminationGoodsOnofflineAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationGoodsOnofflineAPIResponse.Put(v)
}
