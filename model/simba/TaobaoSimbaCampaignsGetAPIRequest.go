package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignsGetAPIRequest 取得一组推广计划 API请求
// taobao.simba.campaigns.get
//
// 取得一个客户的推广计划；
type TaobaoSimbaCampaignsGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 计划类型0位标准计划，16位销量明星计划
	_type int64
}

// NewTaobaoSimbaCampaignsGetRequest 初始化TaobaoSimbaCampaignsGetAPIRequest对象
func NewTaobaoSimbaCampaignsGetRequest() *TaobaoSimbaCampaignsGetAPIRequest {
	return &TaobaoSimbaCampaignsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignsGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaigns.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignsGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCampaignsGetAPIRequest) GetNick() string {
	return r._nick
}

// SetType is Type Setter
// 计划类型0位标准计划，16位销量明星计划
func (r *TaobaoSimbaCampaignsGetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoSimbaCampaignsGetAPIRequest) GetType() int64 {
	return r._type
}
