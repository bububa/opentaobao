package dengta

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPicturesDengtaImsOrderStatusChangeAPIResponse 天下秀回传订单执行状态变动 API返回值
// alibaba.pictures.dengta.ims.order.status.change
//
// 天下秀回传订单执行状态变动
type AlibabaPicturesDengtaImsOrderStatusChangeAPIResponse struct {
	model.CommonResponse
	AlibabaPicturesDengtaImsOrderStatusChangeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPicturesDengtaImsOrderStatusChangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPicturesDengtaImsOrderStatusChangeAPIResponseModel).Reset()
}

// AlibabaPicturesDengtaImsOrderStatusChangeAPIResponseModel is 天下秀回传订单执行状态变动 成功返回结果
type AlibabaPicturesDengtaImsOrderStatusChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pictures_dengta_ims_order_status_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口出参
	Result *GeneralResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPicturesDengtaImsOrderStatusChangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPicturesDengtaImsOrderStatusChangeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPicturesDengtaImsOrderStatusChangeAPIResponse)
	},
}

// GetAlibabaPicturesDengtaImsOrderStatusChangeAPIResponse 从 sync.Pool 获取 AlibabaPicturesDengtaImsOrderStatusChangeAPIResponse
func GetAlibabaPicturesDengtaImsOrderStatusChangeAPIResponse() *AlibabaPicturesDengtaImsOrderStatusChangeAPIResponse {
	return poolAlibabaPicturesDengtaImsOrderStatusChangeAPIResponse.Get().(*AlibabaPicturesDengtaImsOrderStatusChangeAPIResponse)
}

// ReleaseAlibabaPicturesDengtaImsOrderStatusChangeAPIResponse 将 AlibabaPicturesDengtaImsOrderStatusChangeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPicturesDengtaImsOrderStatusChangeAPIResponse(v *AlibabaPicturesDengtaImsOrderStatusChangeAPIResponse) {
	v.Reset()
	poolAlibabaPicturesDengtaImsOrderStatusChangeAPIResponse.Put(v)
}
