package lstspeacker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest
音箱音量调节 API请求
alibaba.lst.speaker.configure.adjustvolume

音箱音量调节 */
type AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 音量直
	_volume string
	// 音量类型，val:固定值, percent:百分比
	_valueType string
}

// New
