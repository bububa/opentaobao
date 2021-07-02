package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaobaoMicdetailAlihealthQuerystoresAPIResponse 疫苗预约门店列表查询 API返回值
// alibaba.taobao.micdetail.alihealth.querystores
//
// 微信小程序疫苗预约门店列表查询
type AlibabaTaobaoMicdetailAlihealthQuerystoresAPIResponse struct {
	model.CommonResponse
	AlibabaTaobaoMicdetailAlihealthQuerystoresAPIResponseModel
}

// AlibabaTaobaoMicdetailAlihealthQuerystoresAPIResponseModel is 疫苗预约门店列表查询 成功返回结果
type AlibabaTaobaoMicdetailAlihealthQuerystoresAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_taobao_micdetail_alihealth_querystores_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaTaobaoMicdetailAlihealthQuerystoresResult `json:"result,omitempty" xml:"result,omitempty"`
}
