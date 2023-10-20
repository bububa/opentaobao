package campus

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaCampusGuardTimeperiodListdatetemplateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusGuardTimeperiodListdatetemplateAPIResponseModel).Reset()
}

// AlibabaCampusGuardTimeperiodListdatetemplateAPIResponseModel is 门禁控制器查询日期模版 成功返回结果
type AlibabaCampusGuardTimeperiodListdatetemplateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guard_timeperiod_listdatetemplate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusGuardTimeperiodListdatetemplateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusGuardTimeperiodListdatetemplateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusGuardTimeperiodListdatetemplateAPIResponse)
	},
}

// GetAlibabaCampusGuardTimeperiodListdatetemplateAPIResponse 从 sync.Pool 获取 AlibabaCampusGuardTimeperiodListdatetemplateAPIResponse
func GetAlibabaCampusGuardTimeperiodListdatetemplateAPIResponse() *AlibabaCampusGuardTimeperiodListdatetemplateAPIResponse {
	return poolAlibabaCampusGuardTimeperiodListdatetemplateAPIResponse.Get().(*AlibabaCampusGuardTimeperiodListdatetemplateAPIResponse)
}

// ReleaseAlibabaCampusGuardTimeperiodListdatetemplateAPIResponse 将 AlibabaCampusGuardTimeperiodListdatetemplateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusGuardTimeperiodListdatetemplateAPIResponse(v *AlibabaCampusGuardTimeperiodListdatetemplateAPIResponse) {
	v.Reset()
	poolAlibabaCampusGuardTimeperiodListdatetemplateAPIResponse.Put(v)
}
