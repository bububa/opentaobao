package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmisccommonitemactivitylistgetAPIRequest 查询通用单品优惠活动列表 API请求
// taobao.promotionmisc.common.item.activity.list.get
//
// 查询通用单品优惠活动列表。
type TaobaopromotionmisccommonitemactivitylistgetAPIRequest struct {
	model.Params
	// 分页页码，页码从1开始
	_pageNo int64
	// 分页大小，不能超过50
	_pageSize int64
}

// NewTaobaopromotionmisccommonitemactivitylistgetRequest 初始化TaobaopromotionmisccommonitemactivitylistgetAPIRequest对象
func NewTaobaopromotionmisccommonitemactivitylistgetRequest() *TaobaopromotionmisccommonitemactivitylistgetAPIRequest {
	return &TaobaopromotionmisccommonitemactivitylistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmisccommonitemactivitylistgetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.activity.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmisccommonitemactivitylistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmisccommonitemactivitylistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageNo is PageNo Setter
// 分页页码，页码从1开始
func (r *TaobaopromotionmisccommonitemactivitylistgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaopromotionmisccommonitemactivitylistgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页大小，不能超过50
func (r *TaobaopromotionmisccommonitemactivitylistgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaopromotionmisccommonitemactivitylistgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
