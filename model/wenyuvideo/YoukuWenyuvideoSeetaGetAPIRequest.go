package wenyuvideo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuWenyuvideoSeetaGetAPIRequest
只看TA API请求
youku.wenyuvideo.seeta.get

只看Ta对外输出 */
type YoukuWenyuvideoSeetaGetAPIRequest struct {
	model.Params
	// 视频字符串形式id
	_videoStrId string
}

// New
