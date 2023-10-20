package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkIotConveyorConveyorconfigGetAPIResponse 获取悬挂链基本配置信息 API返回值
// taobao.wdk.iot.conveyor.conveyorconfig.get
//
// 用于从云端WCS获取悬挂链基本配置信息
type TaobaoWdkIotConveyorConveyorconfigGetAPIResponse struct {
	model.CommonResponse
	TaobaoWdkIotConveyorConveyorconfigGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWdkIotConveyorConveyorconfigGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWdkIotConveyorConveyorconfigGetAPIResponseModel).Reset()
}

// TaobaoWdkIotConveyorConveyorconfigGetAPIResponseModel is 获取悬挂链基本配置信息 成功返回结果
type TaobaoWdkIotConveyorConveyorconfigGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_iot_conveyor_conveyorconfig_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaoWdkIotConveyorConveyorconfigGetHMResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWdkIotConveyorConveyorconfigGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWdkIotConveyorConveyorconfigGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWdkIotConveyorConveyorconfigGetAPIResponse)
	},
}

// GetTaobaoWdkIotConveyorConveyorconfigGetAPIResponse 从 sync.Pool 获取 TaobaoWdkIotConveyorConveyorconfigGetAPIResponse
func GetTaobaoWdkIotConveyorConveyorconfigGetAPIResponse() *TaobaoWdkIotConveyorConveyorconfigGetAPIResponse {
	return poolTaobaoWdkIotConveyorConveyorconfigGetAPIResponse.Get().(*TaobaoWdkIotConveyorConveyorconfigGetAPIResponse)
}

// ReleaseTaobaoWdkIotConveyorConveyorconfigGetAPIResponse 将 TaobaoWdkIotConveyorConveyorconfigGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWdkIotConveyorConveyorconfigGetAPIResponse(v *TaobaoWdkIotConveyorConveyorconfigGetAPIResponse) {
	v.Reset()
	poolTaobaoWdkIotConveyorConveyorconfigGetAPIResponse.Put(v)
}
