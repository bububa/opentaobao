package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmPointAvailableGetAPIRequest CRM会员积分查询开放接口 API请求
// taobao.crm.point.available.get
//
// 查询用户在某个商家的可用积分数
type TaobaoCrmPointAvailableGetAPIRequest struct {
	model.Params
	// 明文nick，可不填，直接填混淆昵称
	_buyerNick string
	// 混淆昵称
	_mixNick string
	// 买家openid
	_openUid string
}

// NewTaobaoCrmPointAvailableGetRequest 初始化TaobaoCrmPointAvailableGetAPIRequest对象
func NewTaobaoCrmPointAvailableGetRequest() *TaobaoCrmPointAvailableGetAPIRequest {
	return &TaobaoCrmPointAvailableGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmPointAvailableGetAPIRequest) Reset() {
	r._buyerNick = ""
	r._mixNick = ""
	r._openUid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmPointAvailableGetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.point.available.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmPointAvailableGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmPointAvailableGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerNick is BuyerNick Setter
// 明文nick，可不填，直接填混淆昵称
func (r *TaobaoCrmPointAvailableGetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoCrmPointAvailableGetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetMixNick is MixNick Setter
// 混淆昵称
func (r *TaobaoCrmPointAvailableGetAPIRequest) SetMixNick(_mixNick string) error {
	r._mixNick = _mixNick
	r.Set("mix_nick", _mixNick)
	return nil
}

// GetMixNick MixNick Getter
func (r TaobaoCrmPointAvailableGetAPIRequest) GetMixNick() string {
	return r._mixNick
}

// SetOpenUid is OpenUid Setter
// 买家openid
func (r *TaobaoCrmPointAvailableGetAPIRequest) SetOpenUid(_openUid string) error {
	r._openUid = _openUid
	r.Set("open_uid", _openUid)
	return nil
}

// GetOpenUid OpenUid Getter
func (r TaobaoCrmPointAvailableGetAPIRequest) GetOpenUid() string {
	return r._openUid
}

var poolTaobaoCrmPointAvailableGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmPointAvailableGetRequest()
	},
}

// GetTaobaoCrmPointAvailableGetRequest 从 sync.Pool 获取 TaobaoCrmPointAvailableGetAPIRequest
func GetTaobaoCrmPointAvailableGetAPIRequest() *TaobaoCrmPointAvailableGetAPIRequest {
	return poolTaobaoCrmPointAvailableGetAPIRequest.Get().(*TaobaoCrmPointAvailableGetAPIRequest)
}

// ReleaseTaobaoCrmPointAvailableGetAPIRequest 将 TaobaoCrmPointAvailableGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmPointAvailableGetAPIRequest(v *TaobaoCrmPointAvailableGetAPIRequest) {
	v.Reset()
	poolTaobaoCrmPointAvailableGetAPIRequest.Put(v)
}
