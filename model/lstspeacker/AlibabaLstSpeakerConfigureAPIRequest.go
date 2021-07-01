package lstspeacker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstSpeakerConfigureAPIRequest
零售通音箱配置通用泛化调用接口 API请求
alibaba.lst.speaker.configure

零售通音箱配置通用泛化调用接口，包括内容、音量、音频等内容 */
type AlibabaLstSpeakerConfigureAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 命令类型setPayTime,adjustVolume，syncAudio，syncAudioAdvert
	_command string
	// 数据体，根据命令不同而不同
	_params string
}

// New
