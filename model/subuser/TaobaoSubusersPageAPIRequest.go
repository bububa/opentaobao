package subuser

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubusersPageAPIRequest 分页获取指定账户的子账号简易信息列表（新isv建议使用taobao.sellercenter.subusers.page接口） API请求
// taobao.subusers.page
//
// 分页获取指定账户的子账号简易信息列表
// （新isv接入建议使用taobao.sellercenter.subusers.page接口）
type TaobaoSubusersPageAPIRequest struct {
	model.Params
	// 主账号用户名
	_userNick string
	// 可以不传，不传的话和老接口返回数据一样。如果传则根据子账号当前状态筛选 1正常   2冻结  3处罚，如果不传默认都返回
	_accountStatus int64
	// 页大小，大于1小于200
	_pageSize int64
	// 页码，大于等于1
	_pageNum int64
}

// NewTaobaoSubusersPageRequest 初始化TaobaoSubusersPageAPIRequest对象
func NewTaobaoSubusersPageRequest() *TaobaoSubusersPageAPIRequest {
	return &TaobaoSubusersPageAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubusersPageAPIRequest) Reset() {
	r._userNick = ""
	r._accountStatus = 0
	r._pageSize = 0
	r._pageNum = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubusersPageAPIRequest) GetApiMethodName() string {
	return "taobao.subusers.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubusersPageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubusersPageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 主账号用户名
func (r *TaobaoSubusersPageAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoSubusersPageAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetAccountStatus is AccountStatus Setter
// 可以不传，不传的话和老接口返回数据一样。如果传则根据子账号当前状态筛选 1正常   2冻结  3处罚，如果不传默认都返回
func (r *TaobaoSubusersPageAPIRequest) SetAccountStatus(_accountStatus int64) error {
	r._accountStatus = _accountStatus
	r.Set("account_status", _accountStatus)
	return nil
}

// GetAccountStatus AccountStatus Getter
func (r TaobaoSubusersPageAPIRequest) GetAccountStatus() int64 {
	return r._accountStatus
}

// SetPageSize is PageSize Setter
// 页大小，大于1小于200
func (r *TaobaoSubusersPageAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSubusersPageAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 页码，大于等于1
func (r *TaobaoSubusersPageAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoSubusersPageAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

var poolTaobaoSubusersPageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubusersPageRequest()
	},
}

// GetTaobaoSubusersPageRequest 从 sync.Pool 获取 TaobaoSubusersPageAPIRequest
func GetTaobaoSubusersPageAPIRequest() *TaobaoSubusersPageAPIRequest {
	return poolTaobaoSubusersPageAPIRequest.Get().(*TaobaoSubusersPageAPIRequest)
}

// ReleaseTaobaoSubusersPageAPIRequest 将 TaobaoSubusersPageAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubusersPageAPIRequest(v *TaobaoSubusersPageAPIRequest) {
	v.Reset()
	poolTaobaoSubusersPageAPIRequest.Put(v)
}
