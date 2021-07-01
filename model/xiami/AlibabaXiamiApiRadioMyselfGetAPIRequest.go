package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiRadioMyselfGetAPIRequest
我的电台 API请求
alibaba.xiami.api.radio.myself.get

我的电台 */
type AlibabaXiamiApiRadioMyselfGetAPIRequest struct {
	model.Params
	// 歌曲数量
	_limit int64
}

// New
