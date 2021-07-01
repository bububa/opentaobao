package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPictureUserinfoGetAPIRequest
查询图片空间用户的信息 API请求
taobao.picture.userinfo.get

查询用户的图片空间使用信息，包括：订购量，已使用容量，免费容量，总的可使用容量，订购有效期，剩余容量 */
type TaobaoPictureUserinfoGetAPIRequest struct {
	model.Params
}

// New
