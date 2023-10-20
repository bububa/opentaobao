package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomethreedimensionsyncAPIResponse 二手房3D户型信息同步 API返回值
// alibaba.alihouse.existinghome.threedimension.sync
//
// 二手房3D户型信息同步
type AlibabaalihouseexistinghomethreedimensionsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomethreedimensionsyncAPIResponseModel
}

// AlibabaalihouseexistinghomethreedimensionsyncAPIResponseModel is 二手房3D户型信息同步 成功返回结果
type AlibabaalihouseexistinghomethreedimensionsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_threedimension_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihouseexistinghomethreedimensionsyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
