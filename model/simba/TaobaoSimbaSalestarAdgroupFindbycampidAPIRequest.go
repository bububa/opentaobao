package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest (销量明星)批量获取推广计划下的推广组信息 API请求
// taobao.simba.salestar.adgroup.findbycampid
//
// 批量得到推广计划下的推广组
type TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest struct {
	model.Params
	// 推广计划Id
	_campaignId int64
	// 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
	_pageSize int64
	// 页码，从1开始
	_pageNo int64
}

// NewTaobaoSimbaSalestarAdgroupFindbycampidRequest 初始化TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest对象
func NewTaobaoSimbaSalestarAdgroupFindbycampidRequest() *TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest {
	return &TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest) Reset() {
	r._campaignId = 0
	r._pageSize = 0
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.adgroup.findbycampid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetPageSize is PageSize Setter
// 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
func (r *TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码，从1开始
func (r *TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolTaobaoSimbaSalestarAdgroupFindbycampidAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaSalestarAdgroupFindbycampidRequest()
	},
}

// GetTaobaoSimbaSalestarAdgroupFindbycampidRequest 从 sync.Pool 获取 TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest
func GetTaobaoSimbaSalestarAdgroupFindbycampidAPIRequest() *TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest {
	return poolTaobaoSimbaSalestarAdgroupFindbycampidAPIRequest.Get().(*TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest)
}

// ReleaseTaobaoSimbaSalestarAdgroupFindbycampidAPIRequest 将 TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaSalestarAdgroupFindbycampidAPIRequest(v *TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest) {
	v.Reset()
	poolTaobaoSimbaSalestarAdgroupFindbycampidAPIRequest.Put(v)
}
