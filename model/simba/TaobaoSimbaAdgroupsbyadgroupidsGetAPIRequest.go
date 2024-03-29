package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest 批量得到推广组 API请求
// taobao.simba.adgroupsbyadgroupids.get
//
// 批量得到推广组
type TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest struct {
	model.Params
	// 推广组Id列表
	_adgroupIds []string
	// 主人昵称
	_nick string
	// 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
	_pageSize int64
	// 页码，从1开始
	_pageNo int64
}

// NewTaobaoSimbaAdgroupsbyadgroupidsGetRequest 初始化TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest对象
func NewTaobaoSimbaAdgroupsbyadgroupidsGetRequest() *TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest {
	return &TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest) Reset() {
	r._adgroupIds = r._adgroupIds[:0]
	r._nick = ""
	r._pageSize = 0
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroupsbyadgroupids.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupIds is AdgroupIds Setter
// 推广组Id列表
func (r *TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest) SetAdgroupIds(_adgroupIds []string) error {
	r._adgroupIds = _adgroupIds
	r.Set("adgroup_ids", _adgroupIds)
	return nil
}

// GetAdgroupIds AdgroupIds Getter
func (r TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest) GetAdgroupIds() []string {
	return r._adgroupIds
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest) GetNick() string {
	return r._nick
}

// SetPageSize is PageSize Setter
// 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
func (r *TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码，从1开始
func (r *TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolTaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaAdgroupsbyadgroupidsGetRequest()
	},
}

// GetTaobaoSimbaAdgroupsbyadgroupidsGetRequest 从 sync.Pool 获取 TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest
func GetTaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest() *TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest {
	return poolTaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest.Get().(*TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest)
}

// ReleaseTaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest 将 TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest(v *TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest.Put(v)
}
