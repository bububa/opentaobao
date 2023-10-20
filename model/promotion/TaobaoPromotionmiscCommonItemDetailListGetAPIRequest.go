package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmisccommonitemdetaillistgetAPIRequest 查询通用单品优惠详情列表 API请求
// taobao.promotionmisc.common.item.detail.list.get
//
// 查询通用单品优惠详情列表。
type TaobaopromotionmisccommonitemdetaillistgetAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
	// 分页页码，页码从1开始
	_pageNo int64
	// 分页大小，不能超过50
	_pageSize int64
}

// NewTaobaopromotionmisccommonitemdetaillistgetRequest 初始化TaobaopromotionmisccommonitemdetaillistgetAPIRequest对象
func NewTaobaopromotionmisccommonitemdetaillistgetRequest() *TaobaopromotionmisccommonitemdetaillistgetAPIRequest {
	return &TaobaopromotionmisccommonitemdetaillistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmisccommonitemdetaillistgetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.detail.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmisccommonitemdetaillistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmisccommonitemdetaillistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 优惠活动ID
func (r *TaobaopromotionmisccommonitemdetaillistgetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaopromotionmisccommonitemdetaillistgetAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetPageNo is PageNo Setter
// 分页页码，页码从1开始
func (r *TaobaopromotionmisccommonitemdetaillistgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaopromotionmisccommonitemdetaillistgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页大小，不能超过50
func (r *TaobaopromotionmisccommonitemdetaillistgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaopromotionmisccommonitemdetaillistgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
