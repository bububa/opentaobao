package smartstore

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPopupstoreActivityDeviceQueryAPIResponse 根据活动id查询活动相关快闪店及设备信息 API返回值
// tmall.popupstore.activity.device.query
//
// 查询某一活动的deviceCode的部署情况
type TmallPopupstoreActivityDeviceQueryAPIResponse struct {
	model.CommonResponse
	TmallPopupstoreActivityDeviceQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallPopupstoreActivityDeviceQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallPopupstoreActivityDeviceQueryAPIResponseModel).Reset()
}

// TmallPopupstoreActivityDeviceQueryAPIResponseModel is 根据活动id查询活动相关快闪店及设备信息 成功返回结果
type TmallPopupstoreActivityDeviceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_popupstore_activity_device_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	ResultDto *TmallPopupstoreActivityDeviceQueryResultDto `json:"result_dto,omitempty" xml:"result_dto,omitempty"`
}

// Reset 清空结构体
func (m *TmallPopupstoreActivityDeviceQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultDto = nil
}

var poolTmallPopupstoreActivityDeviceQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallPopupstoreActivityDeviceQueryAPIResponse)
	},
}

// GetTmallPopupstoreActivityDeviceQueryAPIResponse 从 sync.Pool 获取 TmallPopupstoreActivityDeviceQueryAPIResponse
func GetTmallPopupstoreActivityDeviceQueryAPIResponse() *TmallPopupstoreActivityDeviceQueryAPIResponse {
	return poolTmallPopupstoreActivityDeviceQueryAPIResponse.Get().(*TmallPopupstoreActivityDeviceQueryAPIResponse)
}

// ReleaseTmallPopupstoreActivityDeviceQueryAPIResponse 将 TmallPopupstoreActivityDeviceQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallPopupstoreActivityDeviceQueryAPIResponse(v *TmallPopupstoreActivityDeviceQueryAPIResponse) {
	v.Reset()
	poolTmallPopupstoreActivityDeviceQueryAPIResponse.Put(v)
}
