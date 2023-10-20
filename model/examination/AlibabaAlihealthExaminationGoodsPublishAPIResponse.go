package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationgoodspublishAPIResponse 体检机构对接_商品发布／更新 API返回值
// alibaba.alihealth.examination.goods.publish
//
// 体检机构对接_商品发布／更新
type AlibabaalihealthexaminationgoodspublishAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthexaminationgoodspublishAPIResponseModel
}

// AlibabaalihealthexaminationgoodspublishAPIResponseModel is 体检机构对接_商品发布／更新 成功返回结果
type AlibabaalihealthexaminationgoodspublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_goods_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
