package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaochanneltradeprepayofflinereduceAPIRequest 渠道分销供应商上传线下流水预存款（减少） API请求
// taobao.channel.trade.prepay.offline.reduce
//
// 渠道分销供应商上传线下流水预存款（减少）
type TaobaochanneltradeprepayofflinereduceAPIRequest struct {
	model.Params
	// 减少流水
	_offlineReducePrepayParam *TopOfflineReducePrepayDto
}

// NewTaobaochanneltradeprepayofflinereduceRequest 初始化TaobaochanneltradeprepayofflinereduceAPIRequest对象
func NewTaobaochanneltradeprepayofflinereduceRequest() *TaobaochanneltradeprepayofflinereduceAPIRequest {
	return &TaobaochanneltradeprepayofflinereduceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaochanneltradeprepayofflinereduceAPIRequest) GetApiMethodName() string {
	return "taobao.channel.trade.prepay.offline.reduce"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaochanneltradeprepayofflinereduceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaochanneltradeprepayofflinereduceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineReducePrepayParam is OfflineReducePrepayParam Setter
// 减少流水
func (r *TaobaochanneltradeprepayofflinereduceAPIRequest) SetOfflineReducePrepayParam(_offlineReducePrepayParam *TopOfflineReducePrepayDto) error {
	r._offlineReducePrepayParam = _offlineReducePrepayParam
	r.Set("offline_reduce_prepay_param", _offlineReducePrepayParam)
	return nil
}

// GetOfflineReducePrepayParam OfflineReducePrepayParam Getter
func (r TaobaochanneltradeprepayofflinereduceAPIRequest) GetOfflineReducePrepayParam() *TopOfflineReducePrepayDto {
	return r._offlineReducePrepayParam
}
