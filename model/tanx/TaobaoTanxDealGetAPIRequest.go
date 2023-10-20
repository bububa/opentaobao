package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotanxdealgetAPIRequest 对外部dsp提供交易id查询接口 API请求
// taobao.tanx.deal.get
//
// 对外部dsp提供交易id查询接口
type TaobaotanxdealgetAPIRequest struct {
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

// NewTaobaotanxdealgetRequest 初始化TaobaotanxdealgetAPIRequest对象
func NewTaobaotanxdealgetRequest() *TaobaotanxdealgetAPIRequest {
	return &TaobaotanxdealgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotanxdealgetAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.deal.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotanxdealgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotanxdealgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 验证token
func (r *TaobaotanxdealgetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaotanxdealgetAPIRequest) GetToken() string {
	return r._token
}

// SetDspId is DspId Setter
// dsp用户id
func (r *TaobaotanxdealgetAPIRequest) SetDspId(_dspId int64) error {
	r._dspId = _dspId
	r.Set("dsp_id", _dspId)
	return nil
}

// GetDspId DspId Getter
func (r TaobaotanxdealgetAPIRequest) GetDspId() int64 {
	return r._dspId
}

// SetDealId is DealId Setter
// 交易id
func (r *TaobaotanxdealgetAPIRequest) SetDealId(_dealId int64) error {
	r._dealId = _dealId
	r.Set("deal_id", _dealId)
	return nil
}

// GetDealId DealId Getter
func (r TaobaotanxdealgetAPIRequest) GetDealId() int64 {
	return r._dealId
}

// SetSignTime is SignTime Setter
// 1970年到现在的时间，毫秒
func (r *TaobaotanxdealgetAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaotanxdealgetAPIRequest) GetSignTime() int64 {
	return r._signTime
}
