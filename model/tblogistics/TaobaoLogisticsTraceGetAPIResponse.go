package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticstracegetAPIResponse 用户根据交易号查询物流流转信息 API返回值
// taobao.logistics.trace.get
//
// 用户根据交易号查询物流流转信息
type TaobaologisticstracegetAPIResponse struct {
	model.CommonResponse
	TaobaologisticstracegetAPIResponseModel
}

// TaobaologisticstracegetAPIResponseModel is 用户根据交易号查询物流流转信息 成功返回结果
type TaobaologisticstracegetAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_trace_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result []TransitStepResult `json:"result,omitempty" xml:"result>transit_step_result,omitempty"`
}
