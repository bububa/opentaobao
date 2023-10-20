package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacampaignsgetAPIRequest 取得一组推广计划 API请求
// taobao.simba.campaigns.get
//
// 取得一个客户的推广计划；
type TaobaosimbacampaignsgetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 计划类型0位标准计划，16位销量明星计划
	_type int64
}

// NewTaobaosimbacampaignsgetRequest 初始化TaobaosimbacampaignsgetAPIRequest对象
func NewTaobaosimbacampaignsgetRequest() *TaobaosimbacampaignsgetAPIRequest {
	return &TaobaosimbacampaignsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacampaignsgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaigns.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacampaignsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacampaignsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbacampaignsgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbacampaignsgetAPIRequest) GetNick() string {
	return r._nick
}

// SetType is Type Setter
// 计划类型0位标准计划，16位销量明星计划
func (r *TaobaosimbacampaignsgetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaosimbacampaignsgetAPIRequest) GetType() int64 {
	return r._type
}
