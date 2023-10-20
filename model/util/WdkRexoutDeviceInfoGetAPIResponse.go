package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkRexoutDeviceInfoGetAPIResponse 获取设备详情-外部对接 API返回值
// wdk.rexout.device.info.get
//
// 获取设备详情-外部对接
type WdkRexoutDeviceInfoGetAPIResponse struct {
	model.CommonResponse
	WdkRexoutDeviceInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *WdkRexoutDeviceInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.WdkRexoutDeviceInfoGetAPIResponseModel).Reset()
}

// WdkRexoutDeviceInfoGetAPIResponseModel is 获取设备详情-外部对接 成功返回结果
type WdkRexoutDeviceInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_rexout_device_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 设备详情
	Data *AiotTopOpenDeviceDto `json:"data,omitempty" xml:"data,omitempty"`
	// 结果是否成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}

// Reset 清空结构体
func (m *WdkRexoutDeviceInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Data = nil
	m.Succeed = false
}

var poolWdkRexoutDeviceInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(WdkRexoutDeviceInfoGetAPIResponse)
	},
}

// GetWdkRexoutDeviceInfoGetAPIResponse 从 sync.Pool 获取 WdkRexoutDeviceInfoGetAPIResponse
func GetWdkRexoutDeviceInfoGetAPIResponse() *WdkRexoutDeviceInfoGetAPIResponse {
	return poolWdkRexoutDeviceInfoGetAPIResponse.Get().(*WdkRexoutDeviceInfoGetAPIResponse)
}

// ReleaseWdkRexoutDeviceInfoGetAPIResponse 将 WdkRexoutDeviceInfoGetAPIResponse 保存到 sync.Pool
func ReleaseWdkRexoutDeviceInfoGetAPIResponse(v *WdkRexoutDeviceInfoGetAPIResponse) {
	v.Reset()
	poolWdkRexoutDeviceInfoGetAPIResponse.Put(v)
}
