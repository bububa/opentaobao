package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationGoodsPublishAPIResponse 体检机构对接_商品发布／更新 API返回值
// alibaba.alihealth.examination.goods.publish
//
// 体检机构对接_商品发布／更新
type AlibabaAlihealthExaminationGoodsPublishAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationGoodsPublishAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationGoodsPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationGoodsPublishAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationGoodsPublishAPIResponseModel is 体检机构对接_商品发布／更新 成功返回结果
type AlibabaAlihealthExaminationGoodsPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_goods_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationGoodsPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthExaminationGoodsPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationGoodsPublishAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationGoodsPublishAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationGoodsPublishAPIResponse
func GetAlibabaAlihealthExaminationGoodsPublishAPIResponse() *AlibabaAlihealthExaminationGoodsPublishAPIResponse {
	return poolAlibabaAlihealthExaminationGoodsPublishAPIResponse.Get().(*AlibabaAlihealthExaminationGoodsPublishAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationGoodsPublishAPIResponse 将 AlibabaAlihealthExaminationGoodsPublishAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationGoodsPublishAPIResponse(v *AlibabaAlihealthExaminationGoodsPublishAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationGoodsPublishAPIResponse.Put(v)
}
