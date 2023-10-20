package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardTimeperiodListtimeruleAPIResponse 门禁控制器查询时间规则 API返回值
// alibaba.campus.guard.timeperiod.listtimerule
//
// 门禁控制器查询时间规则
type AlibabaCampusGuardTimeperiodListtimeruleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusGuardTimeperiodListtimeruleAPIResponseModel
}

// AlibabaCampusGuardTimeperiodListtimeruleAPIResponseModel is 门禁控制器查询时间规则 成功返回结果
type AlibabaCampusGuardTimeperiodListtimeruleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guard_timeperiod_listtimerule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
