package ioti

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
相框设备厂测刷图接口 APIResponse
alibaba.it.album.device.sendimage

提供传入电子相框设备mac，mac需属于厂测白名单设备，将设备刷新为系统默认的厂测图片
*/
type AlibabaItAlbumDeviceSendimageAPIResponse struct {
    model.CommonResponse
    AlibabaItAlbumDeviceSendimageResponse
}

type AlibabaItAlbumDeviceSendimageResponse struct {
    XMLName xml.Name `xml:"alibaba_it_album_device_sendimage_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回错误码与参数
    
    Resultmsg   string `json:"resultmsg,omitempty" xml:"resultmsg,omitempty"`

    
}
