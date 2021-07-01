package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWttUserRegioninfoByipGetAPIRequest
根据ip获取省市信息 API请求
alibaba.wtt.user.regioninfo.byip.get

通过ip获取省市信息 */
type AlibabaWttUserRegioninfoByipGetAPIRequest struct {
	model.Params
	// ip地址
	_ip string
}

// New
