package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterservicestoreupdateservicestorecapacityAPIResponse 更新网点容量 API返回值
// tmall.servicecenter.servicestore.updateservicestorecapacity
//
// 更新网点容量，唯一性校验：服务商淘宝账号+网点编码+biz_type
// 前提是网点要存在，
// 如果需要更新的网点容量不存在，会更新失败。
// 网点容量包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、类目区域和容量
type TmallservicecenterservicestoreupdateservicestorecapacityAPIResponse struct {
	model.CommonResponse
	TmallservicecenterservicestoreupdateservicestorecapacityAPIResponseModel
}

// TmallservicecenterservicestoreupdateservicestorecapacityAPIResponseModel is 更新网点容量 成功返回结果
type TmallservicecenterservicestoreupdateservicestorecapacityAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicestore_updateservicestorecapacity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
