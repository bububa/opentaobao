package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoTradePrepayOfflineAddAPIRequest 线下预存款流水增加 API请求
// taobao.fenxiao.trade.prepay.offline.add
//
// 渠道分销供应商上传线下流水预存款（增加）
type TaobaoFenxiaoTradePrepayOfflineAddAPIRequest struct {
	model.Params
	// 增加流水
	_offlineAddPrepayParam *TopOfflineAddPrepayDto
}

// NewTaobaoFenxiaoTradePrepayOfflineAddRequest 初始化TaobaoFenxiaoTradePrepayOfflineAddAPIRequest对象
func NewTaobaoFenxiaoTradePrepayOfflineAddRequest() *TaobaoFenxiaoTradePrepayOfflineAddAPIRequest {
	return &TaobaoFenxiaoTradePrepayOfflineAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoTradePrepayOfflineAddAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.trade.prepay.offline.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoTradePrepayOfflineAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoTradePrepayOfflineAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineAddPrepayParam is OfflineAddPrepayParam Setter
// 增加流水
func (r *TaobaoFenxiaoTradePrepayOfflineAddAPIRequest) SetOfflineAddPrepayParam(_offlineAddPrepayParam *TopOfflineAddPrepayDto) error {
	r._offlineAddPrepayParam = _offlineAddPrepayParam
	r.Set("offline_add_prepay_param", _offlineAddPrepayParam)
	return nil
}

// GetOfflineAddPrepayParam OfflineAddPrepayParam Getter
func (r TaobaoFenxiaoTradePrepayOfflineAddAPIRequest) GetOfflineAddPrepayParam() *TopOfflineAddPrepayDto {
	return r._offlineAddPrepayParam
}
