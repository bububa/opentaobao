package simba

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCampaignsGetAPIRequest) Reset() {
	r._nick = ""
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignsGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaigns.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCampaignsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoSimbaCampaignsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCampaignsGetRequest()
	},
}

// GetTaobaoSimbaCampaignsGetRequest 从 sync.Pool 获取 TaobaoSimbaCampaignsGetAPIRequest
func GetTaobaoSimbaCampaignsGetAPIRequest() *TaobaoSimbaCampaignsGetAPIRequest {
	return poolTaobaoSimbaCampaignsGetAPIRequest.Get().(*TaobaoSimbaCampaignsGetAPIRequest)
}

// ReleaseTaobaoSimbaCampaignsGetAPIRequest 将 TaobaoSimbaCampaignsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCampaignsGetAPIRequest(v *TaobaoSimbaCampaignsGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCampaignsGetAPIRequest.Put(v)
}
