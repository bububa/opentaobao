package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkOrderGetAPIRequest
交易订单详情查询 API请求
alibaba.wdk.order.get

五道口三江单据查询接口 */
type AlibabaWdkOrderGetAPIRequest struct {
	model.Params
	// 系统自动生成
	_idListQueryReq *IdListQueryRequest
}

// New
