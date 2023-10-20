package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmiscitemactivitylistgetAPIRequest 查询无条件单品优惠活动列表 API请求
// taobao.promotionmisc.item.activity.list.get
//
// 查询无条件单品优惠活动列表
type TaobaopromotionmiscitemactivitylistgetAPIRequest struct {
	model.Params
	// 页码。
	_pageNo int64
	// 每页记录数，最大支持50 。
	_pageSize int64
}

// NewTaobaopromotionmiscitemactivitylistgetRequest 初始化TaobaopromotionmiscitemactivitylistgetAPIRequest对象
func NewTaobaopromotionmiscitemactivitylistgetRequest() *TaobaopromotionmiscitemactivitylistgetAPIRequest {
	return &TaobaopromotionmiscitemactivitylistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmiscitemactivitylistgetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.item.activity.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmiscitemactivitylistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmiscitemactivitylistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageNo is PageNo Setter
// 页码。
func (r *TaobaopromotionmiscitemactivitylistgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaopromotionmiscitemactivitylistgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页记录数，最大支持50 。
func (r *TaobaopromotionmiscitemactivitylistgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaopromotionmiscitemactivitylistgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
