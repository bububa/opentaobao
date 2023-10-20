package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusguardantdatasyncAPIResponse 刷卡数据同步 API返回值
// alibaba.campus.guardant.data.sync
//
// 数据同步门禁系统
type AlibabacampusguardantdatasyncAPIResponse struct {
	model.CommonResponse
	AlibabacampusguardantdatasyncAPIResponseModel
}

// AlibabacampusguardantdatasyncAPIResponseModel is 刷卡数据同步 成功返回结果
type AlibabacampusguardantdatasyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guardant_data_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
