package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaotradeprepayofflinereduceAPIRequest 渠道分销供应商上传线下流水预存款（减少） API请求
// taobao.fenxiao.trade.prepay.offline.reduce
//
// 渠道分销供应商上传线下流水预存款（减少）
type TaobaofenxiaotradeprepayofflinereduceAPIRequest struct {
	model.Params
	// 减少流水
	_offlineReducePrepayParam *TopOfflineReducePrepayDto
}

// NewTaobaofenxiaotradeprepayofflinereduceRequest 初始化TaobaofenxiaotradeprepayofflinereduceAPIRequest对象
func NewTaobaofenxiaotradeprepayofflinereduceRequest() *TaobaofenxiaotradeprepayofflinereduceAPIRequest {
	return &TaobaofenxiaotradeprepayofflinereduceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaotradeprepayofflinereduceAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.trade.prepay.offline.reduce"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaotradeprepayofflinereduceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaotradeprepayofflinereduceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineReducePrepayParam is OfflineReducePrepayParam Setter
// 减少流水
func (r *TaobaofenxiaotradeprepayofflinereduceAPIRequest) SetOfflineReducePrepayParam(_offlineReducePrepayParam *TopOfflineReducePrepayDto) error {
	r._offlineReducePrepayParam = _offlineReducePrepayParam
	r.Set("offline_reduce_prepay_param", _offlineReducePrepayParam)
	return nil
}

// GetOfflineReducePrepayParam OfflineReducePrepayParam Getter
func (r TaobaofenxiaotradeprepayofflinereduceAPIRequest) GetOfflineReducePrepayParam() *TopOfflineReducePrepayDto {
	return r._offlineReducePrepayParam
}
