package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignAddAPIRequest 创建一个推广计划 API请求
// taobao.simba.campaign.add
//
// 创建一个推广计划
type TaobaoSimbaCampaignAddAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划名称，不能多余20个汉字，不能和客户其他推广计划同名。
	_title string
	// 计划类型，当前仅支持两种标准推广0，销量明星16，默认为0
	_type int64
}

// NewTaobaoSimbaCampaignAddRequest 初始化TaobaoSimbaCampaignAddAPIRequest对象
func NewTaobaoSimbaCampaignAddRequest() *TaobaoSimbaCampaignAddAPIRequest {
	return &TaobaoSimbaCampaignAddAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCampaignAddAPIRequest) Reset() {
	r._nick = ""
	r._title = ""
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignAddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCampaignAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignAddAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCampaignAddAPIRequest) GetNick() string {
	return r._nick
}

// SetTitle is Title Setter
// 推广计划名称，不能多余20个汉字，不能和客户其他推广计划同名。
func (r *TaobaoSimbaCampaignAddAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoSimbaCampaignAddAPIRequest) GetTitle() string {
	return r._title
}

// SetType is Type Setter
// 计划类型，当前仅支持两种标准推广0，销量明星16，默认为0
func (r *TaobaoSimbaCampaignAddAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoSimbaCampaignAddAPIRequest) GetType() int64 {
	return r._type
}

var poolTaobaoSimbaCampaignAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCampaignAddRequest()
	},
}

// GetTaobaoSimbaCampaignAddRequest 从 sync.Pool 获取 TaobaoSimbaCampaignAddAPIRequest
func GetTaobaoSimbaCampaignAddAPIRequest() *TaobaoSimbaCampaignAddAPIRequest {
	return poolTaobaoSimbaCampaignAddAPIRequest.Get().(*TaobaoSimbaCampaignAddAPIRequest)
}

// ReleaseTaobaoSimbaCampaignAddAPIRequest 将 TaobaoSimbaCampaignAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCampaignAddAPIRequest(v *TaobaoSimbaCampaignAddAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCampaignAddAPIRequest.Put(v)
}
