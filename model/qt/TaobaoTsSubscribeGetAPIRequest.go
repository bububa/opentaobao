package qt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTsSubscribeGetAPIRequest 淘宝服务订购关系查询 API请求
// taobao.ts.subscribe.get
//
// ts订购关系状态查询. 暂只支持1口价服务.
type TaobaoTsSubscribeGetAPIRequest struct {
	model.Params
	// 服务收费项code
	_servcieItemCode string
	// 订购用户昵称
	_nick string
}

// NewTaobaoTsSubscribeGetRequest 初始化TaobaoTsSubscribeGetAPIRequest对象
func NewTaobaoTsSubscribeGetRequest() *TaobaoTsSubscribeGetAPIRequest {
	return &TaobaoTsSubscribeGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTsSubscribeGetAPIRequest) Reset() {
	r._servcieItemCode = ""
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTsSubscribeGetAPIRequest) GetApiMethodName() string {
	return "taobao.ts.subscribe.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTsSubscribeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTsSubscribeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServcieItemCode is ServcieItemCode Setter
// 服务收费项code
func (r *TaobaoTsSubscribeGetAPIRequest) SetServcieItemCode(_servcieItemCode string) error {
	r._servcieItemCode = _servcieItemCode
	r.Set("servcie_item_code", _servcieItemCode)
	return nil
}

// GetServcieItemCode ServcieItemCode Getter
func (r TaobaoTsSubscribeGetAPIRequest) GetServcieItemCode() string {
	return r._servcieItemCode
}

// SetNick is Nick Setter
// 订购用户昵称
func (r *TaobaoTsSubscribeGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoTsSubscribeGetAPIRequest) GetNick() string {
	return r._nick
}

var poolTaobaoTsSubscribeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTsSubscribeGetRequest()
	},
}

// GetTaobaoTsSubscribeGetRequest 从 sync.Pool 获取 TaobaoTsSubscribeGetAPIRequest
func GetTaobaoTsSubscribeGetAPIRequest() *TaobaoTsSubscribeGetAPIRequest {
	return poolTaobaoTsSubscribeGetAPIRequest.Get().(*TaobaoTsSubscribeGetAPIRequest)
}

// ReleaseTaobaoTsSubscribeGetAPIRequest 将 TaobaoTsSubscribeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTsSubscribeGetAPIRequest(v *TaobaoTsSubscribeGetAPIRequest) {
	v.Reset()
	poolTaobaoTsSubscribeGetAPIRequest.Put(v)
}
