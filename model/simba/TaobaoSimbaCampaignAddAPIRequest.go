package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignAddAPIRequest 创建一个推广计划 API请求
// taobao.simba.campaign.add
//
// 创建一个推广计划
type TaobaoSimbaCampaignAddAPIRequest struct {
	model.Params
	// 推广计划名称，不能多余20个汉字，不能和客户其他推广计划同名。
	_title string
	// 主人昵称
	_nick string
	// 计划类型，当前仅支持两种标准推广0，销量明星16，默认为0
	_type int64
}

// NewTaobaoSimbaCampaignAddRequest 初始化TaobaoSimbaCampaignAddAPIRequest对象
func NewTaobaoSimbaCampaignAddRequest() *TaobaoSimbaCampaignAddAPIRequest {
	return &TaobaoSimbaCampaignAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignAddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Title Setter
// 推广计划名称，不能多余20个汉字，不能和客户其他推广计划同名。
func (r *TaobaoSimbaCampaignAddAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// Get Title Getter
func (r TaobaoSimbaCampaignAddAPIRequest) GetTitle() string {
	return r._title
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignAddAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaCampaignAddAPIRequest) GetNick() string {
	return r._nick
}

// Set is Type Setter
// 计划类型，当前仅支持两种标准推广0，销量明星16，默认为0
func (r *TaobaoSimbaCampaignAddAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r TaobaoSimbaCampaignAddAPIRequest) GetType() int64 {
	return r._type
}
