package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRtrptBidwordGetAPIRequest 获取推广词实时报表数据 API请求
// taobao.simba.rtrpt.bidword.get
//
// 获取推广词报表数据
type TaobaoSimbaRtrptBidwordGetAPIRequest struct {
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

// NewTaobaoSimbaRtrptBidwordGetRequest 初始化TaobaoSimbaRtrptBidwordGetAPIRequest对象
func NewTaobaoSimbaRtrptBidwordGetRequest() *TaobaoSimbaRtrptBidwordGetAPIRequest {
	return &TaobaoSimbaRtrptBidwordGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaRtrptBidwordGetAPIRequest) Reset() {
	r._nick = ""
	r._theDate = ""
	r._campaignId = 0
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptBidwordGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rtrpt.bidword.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptBidwordGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaRtrptBidwordGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户名
func (r *TaobaoSimbaRtrptBidwordGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaRtrptBidwordGetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptBidwordGetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaoSimbaRtrptBidwordGetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetCampaignId is CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRtrptBidwordGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaRtrptBidwordGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRtrptBidwordGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaRtrptBidwordGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaRtrptBidwordGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaRtrptBidwordGetRequest()
	},
}

// GetTaobaoSimbaRtrptBidwordGetRequest 从 sync.Pool 获取 TaobaoSimbaRtrptBidwordGetAPIRequest
func GetTaobaoSimbaRtrptBidwordGetAPIRequest() *TaobaoSimbaRtrptBidwordGetAPIRequest {
	return poolTaobaoSimbaRtrptBidwordGetAPIRequest.Get().(*TaobaoSimbaRtrptBidwordGetAPIRequest)
}

// ReleaseTaobaoSimbaRtrptBidwordGetAPIRequest 将 TaobaoSimbaRtrptBidwordGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaRtrptBidwordGetAPIRequest(v *TaobaoSimbaRtrptBidwordGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaRtrptBidwordGetAPIRequest.Put(v)
}
