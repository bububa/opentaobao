package yunosappstore

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosappstoreopenreportadAPIResponse 外投广告上报接口 API返回值
// yunos.appstore.open.reportad
//
// 外投广告回流上报接口
type YunosappstoreopenreportadAPIResponse struct {
	model.CommonResponse
	YunosappstoreopenreportadAPIResponseModel
}

// YunosappstoreopenreportadAPIResponseModel is 外投广告上报接口 成功返回结果
type YunosappstoreopenreportadAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_appstore_open_reportad_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应消息
	Rm string `json:"rm,omitempty" xml:"rm,omitempty"`
	// 响应码
	Rc int64 `json:"rc,omitempty" xml:"rc,omitempty"`
}
