package brandhub

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBrandStartshopRptCampaignGetAPIRequest 明星店铺推广计划报表数据查询 API请求
// taobao.brand.startshop.rpt.campaign.get
//
// 获取明星店铺广告campaign分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaoBrandStartshopRptCampaignGetAPIRequest struct {
	model.Params
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_traffictype string
	// 查询开始时间(最多查询90天数据)
	_startdate string
	// 查询截至时间(最晚查询到昨天)
	_enddate string
	// 转化周期,默认15天,可选 3,7,15
	_effect int64
}

// NewTaobaoBrandStartshopRptCampaignGetRequest 初始化TaobaoBrandStartshopRptCampaignGetAPIRequest对象
func NewTaobaoBrandStartshopRptCampaignGetRequest() *TaobaoBrandStartshopRptCampaignGetAPIRequest {
	return &TaobaoBrandStartshopRptCampaignGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBrandStartshopRptCampaignGetAPIRequest) Reset() {
	r._traffictype = ""
	r._startdate = ""
	r._enddate = ""
	r._effect = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBrandStartshopRptCampaignGetAPIRequest) GetApiMethodName() string {
	return "taobao.brand.startshop.rpt.campaign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBrandStartshopRptCampaignGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBrandStartshopRptCampaignGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraffictype is Traffictype Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoBrandStartshopRptCampaignGetAPIRequest) SetTraffictype(_traffictype string) error {
	r._traffictype = _traffictype
	r.Set("traffictype", _traffictype)
	return nil
}

// GetTraffictype Traffictype Getter
func (r TaobaoBrandStartshopRptCampaignGetAPIRequest) GetTraffictype() string {
	return r._traffictype
}

// SetStartdate is Startdate Setter
// 查询开始时间(最多查询90天数据)
func (r *TaobaoBrandStartshopRptCampaignGetAPIRequest) SetStartdate(_startdate string) error {
	r._startdate = _startdate
	r.Set("startdate", _startdate)
	return nil
}

// GetStartdate Startdate Getter
func (r TaobaoBrandStartshopRptCampaignGetAPIRequest) GetStartdate() string {
	return r._startdate
}

// SetEnddate is Enddate Setter
// 查询截至时间(最晚查询到昨天)
func (r *TaobaoBrandStartshopRptCampaignGetAPIRequest) SetEnddate(_enddate string) error {
	r._enddate = _enddate
	r.Set("enddate", _enddate)
	return nil
}

// GetEnddate Enddate Getter
func (r TaobaoBrandStartshopRptCampaignGetAPIRequest) GetEnddate() string {
	return r._enddate
}

// SetEffect is Effect Setter
// 转化周期,默认15天,可选 3,7,15
func (r *TaobaoBrandStartshopRptCampaignGetAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaoBrandStartshopRptCampaignGetAPIRequest) GetEffect() int64 {
	return r._effect
}

var poolTaobaoBrandStartshopRptCampaignGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBrandStartshopRptCampaignGetRequest()
	},
}

// GetTaobaoBrandStartshopRptCampaignGetRequest 从 sync.Pool 获取 TaobaoBrandStartshopRptCampaignGetAPIRequest
func GetTaobaoBrandStartshopRptCampaignGetAPIRequest() *TaobaoBrandStartshopRptCampaignGetAPIRequest {
	return poolTaobaoBrandStartshopRptCampaignGetAPIRequest.Get().(*TaobaoBrandStartshopRptCampaignGetAPIRequest)
}

// ReleaseTaobaoBrandStartshopRptCampaignGetAPIRequest 将 TaobaoBrandStartshopRptCampaignGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBrandStartshopRptCampaignGetAPIRequest(v *TaobaoBrandStartshopRptCampaignGetAPIRequest) {
	v.Reset()
	poolTaobaoBrandStartshopRptCampaignGetAPIRequest.Put(v)
}
