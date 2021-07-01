package lstspeacker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstSpeakerConfigureSetpaytimeAPIRequest
音箱播放配置 API请求
alibaba.lst.speaker.configure.setpaytime

音箱播放配置 */
type AlibabaLstSpeakerConfigureSetpaytimeAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 开始时间
	_playStartTime string
	// 结束时间
	_playEndTime string
	// 是否播放广告
	_isOnlyPlayAdvert bool
	// 是否设置播放时间
	_isSetPlayTime bool
}

// New
