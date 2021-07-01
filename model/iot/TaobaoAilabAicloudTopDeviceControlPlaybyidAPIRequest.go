package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest
通过id播放歌曲 API请求
taobao.ailab.aicloud.top.device.control.playbyid

通过id播放歌曲 */
type TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// 音频id
	_param2 string
	// 音频来源
	_param3 string
	// 音频类型，如果没有音频类型默认填children_song
	_param4 string
}

// New
