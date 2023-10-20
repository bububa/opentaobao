package ioti

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItAlbumDeviceSendimageAPIResponse 相框设备厂测刷图接口 API返回值
// alibaba.it.album.device.sendimage
//
// 提供传入电子相框设备mac，mac需属于厂测白名单设备，将设备刷新为系统默认的厂测图片
type AlibabaItAlbumDeviceSendimageAPIResponse struct {
	model.CommonResponse
	AlibabaItAlbumDeviceSendimageAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItAlbumDeviceSendimageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItAlbumDeviceSendimageAPIResponseModel).Reset()
}

// AlibabaItAlbumDeviceSendimageAPIResponseModel is 相框设备厂测刷图接口 成功返回结果
type AlibabaItAlbumDeviceSendimageAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_it_album_device_sendimage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回错误码与参数
	Resultmsg string `json:"resultmsg,omitempty" xml:"resultmsg,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItAlbumDeviceSendimageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Resultmsg = ""
}

var poolAlibabaItAlbumDeviceSendimageAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItAlbumDeviceSendimageAPIResponse)
	},
}

// GetAlibabaItAlbumDeviceSendimageAPIResponse 从 sync.Pool 获取 AlibabaItAlbumDeviceSendimageAPIResponse
func GetAlibabaItAlbumDeviceSendimageAPIResponse() *AlibabaItAlbumDeviceSendimageAPIResponse {
	return poolAlibabaItAlbumDeviceSendimageAPIResponse.Get().(*AlibabaItAlbumDeviceSendimageAPIResponse)
}

// ReleaseAlibabaItAlbumDeviceSendimageAPIResponse 将 AlibabaItAlbumDeviceSendimageAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItAlbumDeviceSendimageAPIResponse(v *AlibabaItAlbumDeviceSendimageAPIResponse) {
	v.Reset()
	poolAlibabaItAlbumDeviceSendimageAPIResponse.Put(v)
}
