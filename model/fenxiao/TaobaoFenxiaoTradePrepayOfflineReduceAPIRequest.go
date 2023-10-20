package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest 渠道分销供应商上传线下流水预存款（减少） API请求
// taobao.fenxiao.trade.prepay.offline.reduce
//
// 渠道分销供应商上传线下流水预存款（减少）
type TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest struct {
	model.Params
	// 减少流水
	_offlineReducePrepayParam *TopOfflineReducePrepayDto
}

// NewTaobaoFenxiaoTradePrepayOfflineReduceRequest 初始化TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest对象
func NewTaobaoFenxiaoTradePrepayOfflineReduceRequest() *TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest {
	return &TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) Reset() {
	r._offlineReducePrepayParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.trade.prepay.offline.reduce"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineReducePrepayParam is OfflineReducePrepayParam Setter
// 减少流水
func (r *TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) SetOfflineReducePrepayParam(_offlineReducePrepayParam *TopOfflineReducePrepayDto) error {
	r._offlineReducePrepayParam = _offlineReducePrepayParam
	r.Set("offline_reduce_prepay_param", _offlineReducePrepayParam)
	return nil
}

// GetOfflineReducePrepayParam OfflineReducePrepayParam Getter
func (r TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) GetOfflineReducePrepayParam() *TopOfflineReducePrepayDto {
	return r._offlineReducePrepayParam
}

var poolTaobaoFenxiaoTradePrepayOfflineReduceAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoTradePrepayOfflineReduceRequest()
	},
}

// GetTaobaoFenxiaoTradePrepayOfflineReduceRequest 从 sync.Pool 获取 TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest
func GetTaobaoFenxiaoTradePrepayOfflineReduceAPIRequest() *TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest {
	return poolTaobaoFenxiaoTradePrepayOfflineReduceAPIRequest.Get().(*TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest)
}

// ReleaseTaobaoFenxiaoTradePrepayOfflineReduceAPIRequest 将 TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoTradePrepayOfflineReduceAPIRequest(v *TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoTradePrepayOfflineReduceAPIRequest.Put(v)
}
