package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusGuardDataSyncAPIResponse
卡巴数据同步 API返回值
alibaba.campus.guard.data.sync

数据同步门禁系统 */
type AlibabaCampusGuardDataSyncAPIResponse struct {
	model.CommonResponse
	AlibabaCampusGuardDataSyncAPIResponseModel
}

// AlibabaCampusGuardDataSyncAPIResponseModel is 卡巴数据同步 成功返回结果
type AlibabaCampusGuardDataSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guard_data_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
