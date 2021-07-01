package ascpqcc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpQccSampleCancelItemRelationAPIRequest
魅力惠样品解除父子商品关系 API请求
alibaba.ascp.qcc.sample.cancel.item.relation

品控中心魅力惠样品解除父子商品关系 */
type AlibabaAscpQccSampleCancelItemRelationAPIRequest struct {
	model.Params
	// 请求参数对象
	_cancelRequest *CancelSampleRelationRequest
}

// New
