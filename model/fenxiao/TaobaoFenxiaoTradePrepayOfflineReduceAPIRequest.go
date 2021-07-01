package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest
渠道分销供应商上传线下流水预存款（减少） API请求
taobao.fenxiao.trade.prepay.offline.reduce

渠道分销供应商上传线下流水预存款（减少） */
type TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest struct {
	model.Params
	// 减少流水
	_offlineReducePrepayParam *TopOfflineReducePrepayDto
}

// NewTaobaoFenxiaoTradePrepayOfflineReduceRequest 初始化TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest对象
func NewTaobaoFenxiaoTradePrepayOfflineReduceRequest() *TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest {
	return &TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.trade.prepay.offline.reduce"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OfflineReducePrepayParam Setter
// 减少流水
func (r *TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) SetOfflineReducePrepayParam(_offlineReducePrepayParam *TopOfflineReducePrepayDto) error {
	r._offlineReducePrepayParam = _offlineReducePrepayParam
	r.Set("offline_reduce_prepay_param", _offlineReducePrepayParam)
	return nil
}

// Get OfflineReducePrepayParam Getter
func (r TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) GetOfflineReducePrepayParam() *TopOfflineReducePrepayDto {
	return r._offlineReducePrepayParam
}
