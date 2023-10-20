package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRtrptCreativeGetAPIRequest 获取创意实时报表数据 API请求
// taobao.simba.rtrpt.creative.get
//
// 获取创意实时报表数据
type TaobaoSimbaRtrptCreativeGetAPIRequest struct {
	model.Params
	// 用户名
	_nick string
	// 日期，格式yyyy-mm-dd
	_theDate string
	// 推广计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
}

// NewTaobaoSimbaRtrptCreativeGetRequest 初始化TaobaoSimbaRtrptCreativeGetAPIRequest对象
func NewTaobaoSimbaRtrptCreativeGetRequest() *TaobaoSimbaRtrptCreativeGetAPIRequest {
	return &TaobaoSimbaRtrptCreativeGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaRtrptCreativeGetAPIRequest) Reset() {
	r._nick = ""
	r._theDate = ""
	r._campaignId = 0
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptCreativeGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rtrpt.creative.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptCreativeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaRtrptCreativeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户名
func (r *TaobaoSimbaRtrptCreativeGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaRtrptCreativeGetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptCreativeGetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaoSimbaRtrptCreativeGetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetCampaignId is CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRtrptCreativeGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaRtrptCreativeGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRtrptCreativeGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaRtrptCreativeGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaRtrptCreativeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaRtrptCreativeGetRequest()
	},
}

// GetTaobaoSimbaRtrptCreativeGetRequest 从 sync.Pool 获取 TaobaoSimbaRtrptCreativeGetAPIRequest
func GetTaobaoSimbaRtrptCreativeGetAPIRequest() *TaobaoSimbaRtrptCreativeGetAPIRequest {
	return poolTaobaoSimbaRtrptCreativeGetAPIRequest.Get().(*TaobaoSimbaRtrptCreativeGetAPIRequest)
}

// ReleaseTaobaoSimbaRtrptCreativeGetAPIRequest 将 TaobaoSimbaRtrptCreativeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaRtrptCreativeGetAPIRequest(v *TaobaoSimbaRtrptCreativeGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaRtrptCreativeGetAPIRequest.Put(v)
}
