package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotssubscribegetAPIRequest 淘宝服务订购关系查询 API请求
// taobao.ts.subscribe.get
//
// ts订购关系状态查询. 暂只支持1口价服务.
type TaobaotssubscribegetAPIRequest struct {
	model.Params
	// 服务收费项code
	_servcieItemCode string
	// 订购用户昵称
	_nick string
}

// NewTaobaotssubscribegetRequest 初始化TaobaotssubscribegetAPIRequest对象
func NewTaobaotssubscribegetRequest() *TaobaotssubscribegetAPIRequest {
	return &TaobaotssubscribegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotssubscribegetAPIRequest) GetApiMethodName() string {
	return "taobao.ts.subscribe.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotssubscribegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotssubscribegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServcieItemCode is ServcieItemCode Setter
// 服务收费项code
func (r *TaobaotssubscribegetAPIRequest) SetServcieItemCode(_servcieItemCode string) error {
	r._servcieItemCode = _servcieItemCode
	r.Set("servcie_item_code", _servcieItemCode)
	return nil
}

// GetServcieItemCode ServcieItemCode Getter
func (r TaobaotssubscribegetAPIRequest) GetServcieItemCode() string {
	return r._servcieItemCode
}

// SetNick is Nick Setter
// 订购用户昵称
func (r *TaobaotssubscribegetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaotssubscribegetAPIRequest) GetNick() string {
	return r._nick
}
