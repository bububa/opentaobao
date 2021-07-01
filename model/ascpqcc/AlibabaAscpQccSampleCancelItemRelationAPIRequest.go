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

// NewAlibabaAscpQccSampleCancelItemRelationRequest 初始化AlibabaAscpQccSampleCancelItemRelationAPIRequest对象
func NewAlibabaAscpQccSampleCancelItemRelationRequest() *AlibabaAscpQccSampleCancelItemRelationAPIRequest {
	return &AlibabaAscpQccSampleCancelItemRelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpQccSampleCancelItemRelationAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.qcc.sample.cancel.item.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpQccSampleCancelItemRelationAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CancelRequest Setter
// 请求参数对象
func (r *AlibabaAscpQccSampleCancelItemRelationAPIRequest) SetCancelRequest(_cancelRequest *CancelSampleRelationRequest) error {
	r._cancelRequest = _cancelRequest
	r.Set("cancel_request", _cancelRequest)
	return nil
}

// Get CancelRequest Getter
func (r AlibabaAscpQccSampleCancelItemRelationAPIRequest) GetCancelRequest() *CancelSampleRelationRequest {
	return r._cancelRequest
}
