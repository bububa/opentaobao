package fans

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallfansarenapushAPIRequest 消息推送 API请求
// tmall.fans.arena.push
//
// 超级擂台消息推送
type TmallfansarenapushAPIRequest struct {
	model.Params
	// 推送列表
	_pushList []PushMessageParamDo
}

// NewTmallfansarenapushRequest 初始化TmallfansarenapushAPIRequest对象
func NewTmallfansarenapushRequest() *TmallfansarenapushAPIRequest {
	return &TmallfansarenapushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallfansarenapushAPIRequest) GetApiMethodName() string {
	return "tmall.fans.arena.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallfansarenapushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallfansarenapushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushList is PushList Setter
// 推送列表
func (r *TmallfansarenapushAPIRequest) SetPushList(_pushList []PushMessageParamDo) error {
	r._pushList = _pushList
	r.Set("push_list", _pushList)
	return nil
}

// GetPushList PushList Getter
func (r TmallfansarenapushAPIRequest) GetPushList() []PushMessageParamDo {
	return r._pushList
}
