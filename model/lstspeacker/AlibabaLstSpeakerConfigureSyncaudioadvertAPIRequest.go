package lstspeacker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest
同步广告 API请求
alibaba.lst.speaker.configure.syncaudioadvert

如意音箱广告同步 */
type AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 音频参数
	_speakerConfigParam4SyncAudioAdvert *SpeakerConfigParam4SyncAudioAdvert
}

// New
