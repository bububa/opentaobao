package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpdetailsgetAPIRequest 查询活动详情列表 API请求
// taobao.ump.details.get
//
// 分页查询优惠详情列表
type TaobaoumpdetailsgetAPIRequest struct {
	model.Params
	// 营销活动id
	_actId int64
	// 分页的页码
	_pageNo int64
	// 每页的最大条数
	_pageSize int64
}

// NewTaobaoumpdetailsgetRequest 初始化TaobaoumpdetailsgetAPIRequest对象
func NewTaobaoumpdetailsgetRequest() *TaobaoumpdetailsgetAPIRequest {
	return &TaobaoumpdetailsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumpdetailsgetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.details.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumpdetailsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumpdetailsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActId is ActId Setter
// 营销活动id
func (r *TaobaoumpdetailsgetAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoumpdetailsgetAPIRequest) GetActId() int64 {
	return r._actId
}

// SetPageNo is PageNo Setter
// 分页的页码
func (r *TaobaoumpdetailsgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoumpdetailsgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页的最大条数
func (r *TaobaoumpdetailsgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoumpdetailsgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
