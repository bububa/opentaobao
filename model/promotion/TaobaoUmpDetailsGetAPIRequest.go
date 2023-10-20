package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpDetailsGetAPIRequest 查询活动详情列表 API请求
// taobao.ump.details.get
//
// 分页查询优惠详情列表
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpDetailsGetAPIRequest) Reset() {
	r._actId = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpDetailsGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.details.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpDetailsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpDetailsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActId is ActId Setter
// 营销活动id
func (r *TaobaoUmpDetailsGetAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoUmpDetailsGetAPIRequest) GetActId() int64 {
	return r._actId
}

// SetPageNo is PageNo Setter
// 分页的页码
func (r *TaobaoUmpDetailsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoUmpDetailsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页的最大条数
func (r *TaobaoUmpDetailsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoUmpDetailsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoUmpDetailsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpDetailsGetRequest()
	},
}

// GetTaobaoUmpDetailsGetRequest 从 sync.Pool 获取 TaobaoUmpDetailsGetAPIRequest
func GetTaobaoUmpDetailsGetAPIRequest() *TaobaoUmpDetailsGetAPIRequest {
	return poolTaobaoUmpDetailsGetAPIRequest.Get().(*TaobaoUmpDetailsGetAPIRequest)
}

// ReleaseTaobaoUmpDetailsGetAPIRequest 将 TaobaoUmpDetailsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpDetailsGetAPIRequest(v *TaobaoUmpDetailsGetAPIRequest) {
	v.Reset()
	poolTaobaoUmpDetailsGetAPIRequest.Put(v)
}
