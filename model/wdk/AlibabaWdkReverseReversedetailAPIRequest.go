package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkReverseReversedetailAPIRequest
退款详情 API请求
alibaba.wdk.reverse.reversedetail

退款详情 */
type AlibabaWdkReverseReversedetailAPIRequest struct {
	model.Params
	// 退款单id
	_reverseId string
}

// New
