package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenpushitemAPIRequest 大麦换验平台-第三方对外开放-票品接口pushItem API请求
// alibaba.damai.mev.open.pushitem
//
// 开放接口 推送票品
type AlibabadamaimevopenpushitemAPIRequest struct {
	model.Params
	// 入参pushItemParam
	_pushItemParam *PushTicketItemPushOpenParam
}

// NewAlibabadamaimevopenpushitemRequest 初始化AlibabadamaimevopenpushitemAPIRequest对象
func NewAlibabadamaimevopenpushitemRequest() *AlibabadamaimevopenpushitemAPIRequest {
	return &AlibabadamaimevopenpushitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopenpushitemAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopenpushitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopenpushitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushItemParam is PushItemParam Setter
// 入参pushItemParam
func (r *AlibabadamaimevopenpushitemAPIRequest) SetPushItemParam(_pushItemParam *PushTicketItemPushOpenParam) error {
	r._pushItemParam = _pushItemParam
	r.Set("push_item_param", _pushItemParam)
	return nil
}

// GetPushItemParam PushItemParam Getter
func (r AlibabadamaimevopenpushitemAPIRequest) GetPushItemParam() *PushTicketItemPushOpenParam {
	return r._pushItemParam
}
