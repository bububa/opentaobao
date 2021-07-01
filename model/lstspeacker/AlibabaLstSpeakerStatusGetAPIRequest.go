package lstspeacker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstSpeakerStatusGetAPIRequest
音箱设备在线状态 API请求
alibaba.lst.speaker.status.get

音箱设备在线状态查询 */
type AlibabaLstSpeakerStatusGetAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
}

// New
