package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse 同城零售履约异常中心异常单处理结果回调接口 API返回值
// tmall.cityretail.fulfill.abnormal.center.abnormal.status.change
//
// 同城零售履约异常中心异常单处理结果回调接口
type TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse struct {
	model.CommonResponse
	TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponseModel).Reset()
}

// TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponseModel is 同城零售履约异常中心异常单处理结果回调接口 成功返回结果
type TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_cityretail_fulfill_abnormal_center_abnormal_status_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	FulfillSingleResult *FulfillSingleResult `json:"fulfill_single_result,omitempty" xml:"fulfill_single_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FulfillSingleResult = nil
}

var poolTmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse)
	},
}

// GetTmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse 从 sync.Pool 获取 TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse
func GetTmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse() *TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse {
	return poolTmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse.Get().(*TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse)
}

// ReleaseTmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse 将 TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse 保存到 sync.Pool
func ReleaseTmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse(v *TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse) {
	v.Reset()
	poolTmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse.Put(v)
}
