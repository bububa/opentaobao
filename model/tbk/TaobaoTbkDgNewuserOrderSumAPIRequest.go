package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgnewuserordersumAPIRequest 淘宝客-推广者-拉新活动对应数据查询 API请求
// taobao.tbk.dg.newuser.order.sum
//
// 拉新活动汇总API
type TaobaotbkdgnewuserordersumAPIRequest struct {
	model.Params
	// 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
	_activityId string
	// 结算月份
	_settleMonth string
	// 页大小，默认20，1~100
	_pageSize int64
	// mm_xxx_xxx_xxx的第三位
	_adzoneId int64
	// 页码，默认1
	_pageNo int64
	// mm_xxx_xxx_xxx的第二位
	_siteId int64
}

// NewTaobaotbkdgnewuserordersumRequest 初始化TaobaotbkdgnewuserordersumAPIRequest对象
func NewTaobaotbkdgnewuserordersumRequest() *TaobaotbkdgnewuserordersumAPIRequest {
	return &TaobaotbkdgnewuserordersumAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgnewuserordersumAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.newuser.order.sum"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgnewuserordersumAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgnewuserordersumAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&amp;postId=8599277
func (r *TaobaotbkdgnewuserordersumAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaotbkdgnewuserordersumAPIRequest) GetActivityId() string {
	return r._activityId
}

// SetSettleMonth is SettleMonth Setter
// 结算月份
func (r *TaobaotbkdgnewuserordersumAPIRequest) SetSettleMonth(_settleMonth string) error {
	r._settleMonth = _settleMonth
	r.Set("settle_month", _settleMonth)
	return nil
}

// GetSettleMonth SettleMonth Getter
func (r TaobaotbkdgnewuserordersumAPIRequest) GetSettleMonth() string {
	return r._settleMonth
}

// SetPageSize is PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaotbkdgnewuserordersumAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaotbkdgnewuserordersumAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_xxx的第三位
func (r *TaobaotbkdgnewuserordersumAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaotbkdgnewuserordersumAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetPageNo is PageNo Setter
// 页码，默认1
func (r *TaobaotbkdgnewuserordersumAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaotbkdgnewuserordersumAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetSiteId is SiteId Setter
// mm_xxx_xxx_xxx的第二位
func (r *TaobaotbkdgnewuserordersumAPIRequest) SetSiteId(_siteId int64) error {
	r._siteId = _siteId
	r.Set("site_id", _siteId)
	return nil
}

// GetSiteId SiteId Getter
func (r TaobaotbkdgnewuserordersumAPIRequest) GetSiteId() int64 {
	return r._siteId
}
