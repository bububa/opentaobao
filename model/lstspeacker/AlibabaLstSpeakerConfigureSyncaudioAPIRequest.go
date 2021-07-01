package lstspeacker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstSpeakerConfigureSyncaudioAPIRequest
音频同步 API请求
alibaba.lst.speaker.configure.syncaudio

音频同步 */
type AlibabaLstSpeakerConfigureSyncaudioAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 参数
	_speakerConfigParam4SyncAudio *SpeakerConfigParam4SyncAudio
}

// New
