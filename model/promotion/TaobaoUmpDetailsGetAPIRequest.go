package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpDetailsGetAPIRequest
查询活动详情列表 API请求
taobao.ump.details.get

分页查询优惠详情列表 */
type TaobaoUmpDetailsGetAPIRequest struct {
	model.Params
	// 营销活动id
	_actId int64
	// 分页的页码
	_pageNo int64
	// 每页的最大条数
	_pageSize int64
}

// NewTaobaoUmpDetailsGetRequest 初始化TaobaoUmpDetailsGetAPIRequest对象
func NewTaobaoUmpDetailsGetRequest() *TaobaoUmpDetailsGetAPIRequest {
	return &TaobaoUmpDetailsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpDetailsGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.details.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpDetailsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActId Setter
// 营销活动id
func (r *TaobaoUmpDetailsGetAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// Get ActId Getter
func (r TaobaoUmpDetailsGetAPIRequest) GetActId() int64 {
	return r._actId
}

// Set is PageNo Setter
// 分页的页码
func (r *TaobaoUmpDetailsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoUmpDetailsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 每页的最大条数
func (r *TaobaoUmpDetailsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoUmpDetailsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
