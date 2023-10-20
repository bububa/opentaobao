package ioti

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitalbumdevicesendimageAPIResponse 相框设备厂测刷图接口 API返回值
// alibaba.it.album.device.sendimage
//
// 提供传入电子相框设备mac，mac需属于厂测白名单设备，将设备刷新为系统默认的厂测图片
type AlibabaitalbumdevicesendimageAPIResponse struct {
	model.CommonResponse
	AlibabaitalbumdevicesendimageAPIResponseModel
}

// AlibabaitalbumdevicesendimageAPIResponseModel is 相框设备厂测刷图接口 成功返回结果
type AlibabaitalbumdevicesendimageAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_it_album_device_sendimage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回错误码与参数
	Resultmsg string `json:"resultmsg,omitempty" xml:"resultmsg,omitempty"`
}
