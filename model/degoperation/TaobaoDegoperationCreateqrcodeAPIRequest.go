package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDegoperationCreateqrcodeAPIRequest
中奖生成二维码 API请求
taobao.degoperation.createqrcode

用户中奖后，生成二维码图片链接 */
type TaobaoDegoperationCreateqrcodeAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
	// 系统信息
	_degAccessToken string
	// 奖品唯一标识
	_sequenceNo string
	// 活动名称
	_activity string
	// 奖品名称
	_title string
}

// New
