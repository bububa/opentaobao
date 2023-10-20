package homeai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoHomeaiAlgPredictAPIResponse 硬装预测接口 API返回值
// taobao.homeai.alg.predict
//
// 居然之家硬装预测服务
type TaobaoHomeaiAlgPredictAPIResponse struct {
	model.CommonResponse
	TaobaoHomeaiAlgPredictAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoHomeaiAlgPredictAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoHomeaiAlgPredictAPIResponseModel).Reset()
}

// TaobaoHomeaiAlgPredictAPIResponseModel is 硬装预测接口 成功返回结果
type TaobaoHomeaiAlgPredictAPIResponseModel struct {
	XMLName xml.Name `xml:"homeai_alg_predict_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaoHomeaiAlgPredictResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoHomeaiAlgPredictAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoHomeaiAlgPredictAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoHomeaiAlgPredictAPIResponse)
	},
}

// GetTaobaoHomeaiAlgPredictAPIResponse 从 sync.Pool 获取 TaobaoHomeaiAlgPredictAPIResponse
func GetTaobaoHomeaiAlgPredictAPIResponse() *TaobaoHomeaiAlgPredictAPIResponse {
	return poolTaobaoHomeaiAlgPredictAPIResponse.Get().(*TaobaoHomeaiAlgPredictAPIResponse)
}

// ReleaseTaobaoHomeaiAlgPredictAPIResponse 将 TaobaoHomeaiAlgPredictAPIResponse 保存到 sync.Pool
func ReleaseTaobaoHomeaiAlgPredictAPIResponse(v *TaobaoHomeaiAlgPredictAPIResponse) {
	v.Reset()
	poolTaobaoHomeaiAlgPredictAPIResponse.Put(v)
}
