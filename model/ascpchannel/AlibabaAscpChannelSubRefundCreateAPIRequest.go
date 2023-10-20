package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelsubrefundcreateAPIRequest 淘外分销逆向创单（子单退） API请求
// alibaba.ascp.channel.sub.refund.create
//
// 淘外分销逆向创单（子单退）
type AlibabaascpchannelsubrefundcreateAPIRequest struct {
	model.Params
	// 子单退款创建请求
	_subRefundCreateReq *ExternalCreateRefundOrderDetailRequest
}

// NewAlibabaascpchannelsubrefundcreateRequest 初始化AlibabaascpchannelsubrefundcreateAPIRequest对象
func NewAlibabaascpchannelsubrefundcreateRequest() *AlibabaascpchannelsubrefundcreateAPIRequest {
	return &AlibabaascpchannelsubrefundcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchannelsubrefundcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.sub.refund.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchannelsubrefundcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchannelsubrefundcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubRefundCreateReq is SubRefundCreateReq Setter
// 子单退款创建请求
func (r *AlibabaascpchannelsubrefundcreateAPIRequest) SetSubRefundCreateReq(_subRefundCreateReq *ExternalCreateRefundOrderDetailRequest) error {
	r._subRefundCreateReq = _subRefundCreateReq
	r.Set("sub_refund_create_req", _subRefundCreateReq)
	return nil
}

// GetSubRefundCreateReq SubRefundCreateReq Getter
func (r AlibabaascpchannelsubrefundcreateAPIRequest) GetSubRefundCreateReq() *ExternalCreateRefundOrderDetailRequest {
	return r._subRefundCreateReq
}
