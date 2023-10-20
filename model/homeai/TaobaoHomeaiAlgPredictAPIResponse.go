package homeai

import (
	"encoding/xml"

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

// TaobaoHomeaiAlgPredictAPIResponseModel is 硬装预测接口 成功返回结果
type TaobaoHomeaiAlgPredictAPIResponseModel struct {
	XMLName xml.Name `xml:"homeai_alg_predict_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaoHomeaiAlgPredictResult `json:"result,omitempty" xml:"result,omitempty"`
}
