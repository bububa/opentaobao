package homeai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
硬装预测接口 APIResponse
taobao.homeai.alg.predict

居然之家硬装预测服务
*/
type TaobaoHomeaiAlgPredictAPIResponse struct {
    model.CommonResponse
    TaobaoHomeaiAlgPredictResponse
}

type TaobaoHomeaiAlgPredictResponse struct {
    XMLName xml.Name `xml:"homeai_alg_predict_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值
    
    Result   *TaobaoHomeaiAlgPredictResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
