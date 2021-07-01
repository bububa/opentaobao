package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtStallPayratioSynchronizeAPIRequest
同步摊位收银比例 API请求
tmall.nrt.stall.payratio.synchronize

ISV同步摊位收银比例到阿里 */
type TmallNrtStallPayratioSynchronizeAPIRequest struct {
	model.Params
	// 业务编码
	_bizCode string
	// 合同编号
	_contractCode string
	// 摊位编码
	_storeCode string
	// 收银比例
	_payRatio string
}

// New
