package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthCodesellerGetuserawardAPIRequest
贩卖机扫码查询领奖状态 API请求
alibaba.alihealth.codeseller.getuseraward

贩卖机扫码查询领奖状态 */
type AlibabaAlihealthCodesellerGetuserawardAPIRequest struct {
	model.Params
	// 追溯码
	_code string
}

// New
