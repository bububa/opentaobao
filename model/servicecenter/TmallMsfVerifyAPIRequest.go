package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMsfVerifyAPIRequest
喵师傅核销接口 API请求
tmall.msf.verify

msf服务核销的top接口 */
type TmallMsfVerifyAPIRequest struct {
	model.Params
	// 111
	_shopId string
	// 111
	_bizType string
	// 111
	_code string
}

// New
