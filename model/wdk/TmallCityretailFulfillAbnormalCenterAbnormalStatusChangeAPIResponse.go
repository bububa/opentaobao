package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse
同城零售履约异常中心异常单处理结果回调接口 API返回值
tmall.cityretail.fulfill.abnormal.center.abnormal.status.change

同城零售履约异常中心异常单处理结果回调接口 */
type TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse struct {
	model.CommonResponse
	TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponseModel
}

// TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponseModel is 同城零售履约异常中心异常单处理结果回调接口 成功返回结果
type TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_cityretail_fulfill_abnormal_center_abnormal_status_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	FulfillSingleResult *FulfillSingleResult `json:"fulfill_single_result,omitempty" xml:"fulfill_single_result,omitempty"`
}
