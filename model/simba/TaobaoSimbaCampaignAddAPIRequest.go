package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacampaignaddAPIRequest 创建一个推广计划 API请求
// taobao.simba.campaign.add
//
// 创建一个推广计划
type TaobaosimbacampaignaddAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划名称，不能多余20个汉字，不能和客户其他推广计划同名。
	_title string
	// 计划类型，当前仅支持两种标准推广0，销量明星16，默认为0
	_type int64
}

// NewTaobaosimbacampaignaddRequest 初始化TaobaosimbacampaignaddAPIRequest对象
func NewTaobaosimbacampaignaddRequest() *TaobaosimbacampaignaddAPIRequest {
	return &TaobaosimbacampaignaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacampaignaddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacampaignaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacampaignaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbacampaignaddAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbacampaignaddAPIRequest) GetNick() string {
	return r._nick
}

// SetTitle is Title Setter
// 推广计划名称，不能多余20个汉字，不能和客户其他推广计划同名。
func (r *TaobaosimbacampaignaddAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaosimbacampaignaddAPIRequest) GetTitle() string {
	return r._title
}

// SetType is Type Setter
// 计划类型，当前仅支持两种标准推广0，销量明星16，默认为0
func (r *TaobaosimbacampaignaddAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaosimbacampaignaddAPIRequest) GetType() int64 {
	return r._type
}
