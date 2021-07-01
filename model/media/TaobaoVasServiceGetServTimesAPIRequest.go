package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVasServiceGetServTimesAPIRequest
查询某个用户图片空间的使用情况 API请求
taobao.vas.service.getServTimes

查询某个用户图片空间的使用情况 */
type TaobaoVasServiceGetServTimesAPIRequest struct {
	model.Params
	// 服务编码
	_servCode string
}

// New
