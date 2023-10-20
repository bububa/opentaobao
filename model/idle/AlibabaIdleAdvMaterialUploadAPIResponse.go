package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleadvmaterialuploadAPIResponse 闲鱼用户增长素材中心素材上传接口 API返回值
// alibaba.idle.adv.material.upload
//
// 闲鱼用户增长素材中心素材上传接口
type AlibabaidleadvmaterialuploadAPIResponse struct {
	model.CommonResponse
	AlibabaidleadvmaterialuploadAPIResponseModel
}

// AlibabaidleadvmaterialuploadAPIResponseModel is 闲鱼用户增长素材中心素材上传接口 成功返回结果
type AlibabaidleadvmaterialuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_adv_material_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *IdleAdvResult `json:"result,omitempty" xml:"result,omitempty"`
}
