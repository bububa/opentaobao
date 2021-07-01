package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMsfReceiveAPIRequest
签收接口 API请求
tmall.msf.receive

签收接口 */
type TmallMsfReceiveAPIRequest struct {
	model.Params
	// 1
	_shopId string
	// 1
	_bizType string
	// 1
	_code string
}

// New
