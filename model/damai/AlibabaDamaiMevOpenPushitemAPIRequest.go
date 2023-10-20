package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushitemAPIRequest 大麦换验平台-第三方对外开放-票品接口pushItem API请求
// alibaba.damai.mev.open.pushitem
//
// 开放接口 推送票品
type AlibabaDamaiMevOpenPushitemAPIRequest struct {
	model.Params
	// 入参pushItemParam
	_pushItemParam *PushTicketItemPushOpenParam
}

// NewAlibabaDamaiMevOpenPushitemRequest 初始化AlibabaDamaiMevOpenPushitemAPIRequest对象
func NewAlibabaDamaiMevOpenPushitemRequest() *AlibabaDamaiMevOpenPushitemAPIRequest {
	return &AlibabaDamaiMevOpenPushitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushitemAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenPushitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushItemParam is PushItemParam Setter
// 入参pushItemParam
func (r *AlibabaDamaiMevOpenPushitemAPIRequest) SetPushItemParam(_pushItemParam *PushTicketItemPushOpenParam) error {
	r._pushItemParam = _pushItemParam
	r.Set("push_item_param", _pushItemParam)
	return nil
}

// GetPushItemParam PushItemParam Getter
func (r AlibabaDamaiMevOpenPushitemAPIRequest) GetPushItemParam() *PushTicketItemPushOpenParam {
	return r._pushItemParam
}
