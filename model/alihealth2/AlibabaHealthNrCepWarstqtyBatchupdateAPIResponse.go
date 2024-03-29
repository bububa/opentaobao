package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthNrCepWarstqtyBatchupdateAPIResponse 批量更新ISV库存 API返回值
// alibaba.health.nr.cep.warstqty.batchupdate
//
// 青岛医保服务-ISV批量更新孔雀翎中库存数据
type AlibabaHealthNrCepWarstqtyBatchupdateAPIResponse struct {
	model.CommonResponse
	AlibabaHealthNrCepWarstqtyBatchupdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthNrCepWarstqtyBatchupdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthNrCepWarstqtyBatchupdateAPIResponseModel).Reset()
}

// AlibabaHealthNrCepWarstqtyBatchupdateAPIResponseModel is 批量更新ISV库存 成功返回结果
type AlibabaHealthNrCepWarstqtyBatchupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_nr_cep_warstqty_batchupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *ResponseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthNrCepWarstqtyBatchupdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHealthNrCepWarstqtyBatchupdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthNrCepWarstqtyBatchupdateAPIResponse)
	},
}

// GetAlibabaHealthNrCepWarstqtyBatchupdateAPIResponse 从 sync.Pool 获取 AlibabaHealthNrCepWarstqtyBatchupdateAPIResponse
func GetAlibabaHealthNrCepWarstqtyBatchupdateAPIResponse() *AlibabaHealthNrCepWarstqtyBatchupdateAPIResponse {
	return poolAlibabaHealthNrCepWarstqtyBatchupdateAPIResponse.Get().(*AlibabaHealthNrCepWarstqtyBatchupdateAPIResponse)
}

// ReleaseAlibabaHealthNrCepWarstqtyBatchupdateAPIResponse 将 AlibabaHealthNrCepWarstqtyBatchupdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthNrCepWarstqtyBatchupdateAPIResponse(v *AlibabaHealthNrCepWarstqtyBatchupdateAPIResponse) {
	v.Reset()
	poolAlibabaHealthNrCepWarstqtyBatchupdateAPIResponse.Put(v)
}
