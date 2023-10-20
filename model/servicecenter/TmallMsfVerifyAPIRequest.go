package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMsfVerifyAPIRequest 喵师傅核销接口 API请求
// tmall.msf.verify
//
// msf服务核销的top接口
type TmallMsfVerifyAPIRequest struct {
	model.Params
	// 111
	_shopId string
	// 111
	_bizType string
	// 111
	_code string
}

// NewTmallMsfVerifyRequest 初始化TmallMsfVerifyAPIRequest对象
func NewTmallMsfVerifyRequest() *TmallMsfVerifyAPIRequest {
	return &TmallMsfVerifyAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallMsfVerifyAPIRequest) Reset() {
	r._shopId = ""
	r._bizType = ""
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMsfVerifyAPIRequest) GetApiMethodName() string {
	return "tmall.msf.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMsfVerifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallMsfVerifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopId is ShopId Setter
// 111
func (r *TmallMsfVerifyAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TmallMsfVerifyAPIRequest) GetShopId() string {
	return r._shopId
}

// SetBizType is BizType Setter
// 111
func (r *TmallMsfVerifyAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallMsfVerifyAPIRequest) GetBizType() string {
	return r._bizType
}

// SetCode is Code Setter
// 111
func (r *TmallMsfVerifyAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TmallMsfVerifyAPIRequest) GetCode() string {
	return r._code
}

var poolTmallMsfVerifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallMsfVerifyRequest()
	},
}

// GetTmallMsfVerifyRequest 从 sync.Pool 获取 TmallMsfVerifyAPIRequest
func GetTmallMsfVerifyAPIRequest() *TmallMsfVerifyAPIRequest {
	return poolTmallMsfVerifyAPIRequest.Get().(*TmallMsfVerifyAPIRequest)
}

// ReleaseTmallMsfVerifyAPIRequest 将 TmallMsfVerifyAPIRequest 放入 sync.Pool
func ReleaseTmallMsfVerifyAPIRequest(v *TmallMsfVerifyAPIRequest) {
	v.Reset()
	poolTmallMsfVerifyAPIRequest.Put(v)
}
