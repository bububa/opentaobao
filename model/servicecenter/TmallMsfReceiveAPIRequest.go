package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMsfReceiveAPIRequest 签收接口 API请求
// tmall.msf.receive
//
// 签收接口
type TmallMsfReceiveAPIRequest struct {
	model.Params
	// 1
	_shopId string
	// 1
	_bizType string
	// 1
	_code string
}

// NewTmallMsfReceiveRequest 初始化TmallMsfReceiveAPIRequest对象
func NewTmallMsfReceiveRequest() *TmallMsfReceiveAPIRequest {
	return &TmallMsfReceiveAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallMsfReceiveAPIRequest) Reset() {
	r._shopId = ""
	r._bizType = ""
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMsfReceiveAPIRequest) GetApiMethodName() string {
	return "tmall.msf.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMsfReceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallMsfReceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopId is ShopId Setter
// 1
func (r *TmallMsfReceiveAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TmallMsfReceiveAPIRequest) GetShopId() string {
	return r._shopId
}

// SetBizType is BizType Setter
// 1
func (r *TmallMsfReceiveAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallMsfReceiveAPIRequest) GetBizType() string {
	return r._bizType
}

// SetCode is Code Setter
// 1
func (r *TmallMsfReceiveAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TmallMsfReceiveAPIRequest) GetCode() string {
	return r._code
}

var poolTmallMsfReceiveAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallMsfReceiveRequest()
	},
}

// GetTmallMsfReceiveRequest 从 sync.Pool 获取 TmallMsfReceiveAPIRequest
func GetTmallMsfReceiveAPIRequest() *TmallMsfReceiveAPIRequest {
	return poolTmallMsfReceiveAPIRequest.Get().(*TmallMsfReceiveAPIRequest)
}

// ReleaseTmallMsfReceiveAPIRequest 将 TmallMsfReceiveAPIRequest 放入 sync.Pool
func ReleaseTmallMsfReceiveAPIRequest(v *TmallMsfReceiveAPIRequest) {
	v.Reset()
	poolTmallMsfReceiveAPIRequest.Put(v)
}
