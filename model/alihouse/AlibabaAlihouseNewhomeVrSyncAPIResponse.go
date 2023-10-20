package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomevrsyncAPIResponse VR关系数据同步 API返回值
// alibaba.alihouse.newhome.vr.sync
//
// 对接易居VR关系数据迁移
type AlibabaalihousenewhomevrsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomevrsyncAPIResponseModel
}

// AlibabaalihousenewhomevrsyncAPIResponseModel is VR关系数据同步 成功返回结果
type AlibabaalihousenewhomevrsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_vr_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomevrsyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
