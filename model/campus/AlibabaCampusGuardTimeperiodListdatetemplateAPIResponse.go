package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardTimeperiodListdatetemplateAPIResponse 门禁控制器查询日期模版 API返回值
// alibaba.campus.guard.timeperiod.listdatetemplate
//
// 门禁控制器查询日期模版
type AlibabaCampusGuardTimeperiodListdatetemplateAPIResponse struct {
	model.CommonResponse
	AlibabaCampusGuardTimeperiodListdatetemplateAPIResponseModel
}

// AlibabaCampusGuardTimeperiodListdatetemplateAPIResponseModel is 门禁控制器查询日期模版 成功返回结果
type AlibabaCampusGuardTimeperiodListdatetemplateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guard_timeperiod_listdatetemplate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
