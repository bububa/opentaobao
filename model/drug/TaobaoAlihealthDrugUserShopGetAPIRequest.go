package drug

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlihealthDrugUserShopGetAPIRequest 根据用户id获取店铺id API请求
// taobao.alihealth.drug.user.shop.get
//
// 提供给千牛智能客服，获取用户当前咨询的店铺ID
type TaobaoAlihealthDrugUserShopGetAPIRequest struct {
	model.Params
	// 用户昵称
	_userNick string
}

// NewTaobaoAlihealthDrugUserShopGetRequest 初始化TaobaoAlihealthDrugUserShopGetAPIRequest对象
func NewTaobaoAlihealthDrugUserShopGetRequest() *TaobaoAlihealthDrugUserShopGetAPIRequest {
	return &TaobaoAlihealthDrugUserShopGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlihealthDrugUserShopGetAPIRequest) Reset() {
	r._userNick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlihealthDrugUserShopGetAPIRequest) GetApiMethodName() string {
	return "taobao.alihealth.drug.user.shop.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlihealthDrugUserShopGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlihealthDrugUserShopGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 用户昵称
func (r *TaobaoAlihealthDrugUserShopGetAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoAlihealthDrugUserShopGetAPIRequest) GetUserNick() string {
	return r._userNick
}

var poolTaobaoAlihealthDrugUserShopGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlihealthDrugUserShopGetRequest()
	},
}

// GetTaobaoAlihealthDrugUserShopGetRequest 从 sync.Pool 获取 TaobaoAlihealthDrugUserShopGetAPIRequest
func GetTaobaoAlihealthDrugUserShopGetAPIRequest() *TaobaoAlihealthDrugUserShopGetAPIRequest {
	return poolTaobaoAlihealthDrugUserShopGetAPIRequest.Get().(*TaobaoAlihealthDrugUserShopGetAPIRequest)
}

// ReleaseTaobaoAlihealthDrugUserShopGetAPIRequest 将 TaobaoAlihealthDrugUserShopGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlihealthDrugUserShopGetAPIRequest(v *TaobaoAlihealthDrugUserShopGetAPIRequest) {
	v.Reset()
	poolTaobaoAlihealthDrugUserShopGetAPIRequest.Put(v)
}
