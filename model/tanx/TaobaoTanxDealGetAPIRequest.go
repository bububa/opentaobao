package tanx

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxDealGetAPIRequest 对外部dsp提供交易id查询接口 API请求
// taobao.tanx.deal.get
//
// 对外部dsp提供交易id查询接口
type TaobaoTanxDealGetAPIRequest struct {
	model.Params
	// 验证token
	_token string
	// dsp用户id
	_dspId int64
	// 交易id
	_dealId int64
	// 1970年到现在的时间，毫秒
	_signTime int64
}

// NewTaobaoTanxDealGetRequest 初始化TaobaoTanxDealGetAPIRequest对象
func NewTaobaoTanxDealGetRequest() *TaobaoTanxDealGetAPIRequest {
	return &TaobaoTanxDealGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTanxDealGetAPIRequest) Reset() {
	r._token = ""
	r._dspId = 0
	r._dealId = 0
	r._signTime = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxDealGetAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.deal.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxDealGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTanxDealGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 验证token
func (r *TaobaoTanxDealGetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoTanxDealGetAPIRequest) GetToken() string {
	return r._token
}

// SetDspId is DspId Setter
// dsp用户id
func (r *TaobaoTanxDealGetAPIRequest) SetDspId(_dspId int64) error {
	r._dspId = _dspId
	r.Set("dsp_id", _dspId)
	return nil
}

// GetDspId DspId Getter
func (r TaobaoTanxDealGetAPIRequest) GetDspId() int64 {
	return r._dspId
}

// SetDealId is DealId Setter
// 交易id
func (r *TaobaoTanxDealGetAPIRequest) SetDealId(_dealId int64) error {
	r._dealId = _dealId
	r.Set("deal_id", _dealId)
	return nil
}

// GetDealId DealId Getter
func (r TaobaoTanxDealGetAPIRequest) GetDealId() int64 {
	return r._dealId
}

// SetSignTime is SignTime Setter
// 1970年到现在的时间，毫秒
func (r *TaobaoTanxDealGetAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaoTanxDealGetAPIRequest) GetSignTime() int64 {
	return r._signTime
}

var poolTaobaoTanxDealGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTanxDealGetRequest()
	},
}

// GetTaobaoTanxDealGetRequest 从 sync.Pool 获取 TaobaoTanxDealGetAPIRequest
func GetTaobaoTanxDealGetAPIRequest() *TaobaoTanxDealGetAPIRequest {
	return poolTaobaoTanxDealGetAPIRequest.Get().(*TaobaoTanxDealGetAPIRequest)
}

// ReleaseTaobaoTanxDealGetAPIRequest 将 TaobaoTanxDealGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTanxDealGetAPIRequest(v *TaobaoTanxDealGetAPIRequest) {
	v.Reset()
	poolTaobaoTanxDealGetAPIRequest.Put(v)
}
