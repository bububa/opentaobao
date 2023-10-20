package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosellercentersubuserspageAPIRequest 通过主账号登陆态分页查询子账号列表 API请求
// taobao.sellercenter.subusers.page
//
// 通过主账号登陆态分页查询子账号列表
type TaobaosellercentersubuserspageAPIRequest struct {
	model.Params
	// 主账号登陆nick
	_nick string
	// 可以不传，不传的话和老接口返回数据一样。如果传则根据子账号当前状态筛选 1正常   2冻结  3处罚，如果不传默认都返回
	_accountStatus int64
	// 页大小，大于1小于200
	_pageSize int64
	// 页码，大于等于1
	_pageNum int64
}

// NewTaobaosellercentersubuserspageRequest 初始化TaobaosellercentersubuserspageAPIRequest对象
func NewTaobaosellercentersubuserspageRequest() *TaobaosellercentersubuserspageAPIRequest {
	return &TaobaosellercentersubuserspageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosellercentersubuserspageAPIRequest) GetApiMethodName() string {
	return "taobao.sellercenter.subusers.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosellercentersubuserspageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosellercentersubuserspageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主账号登陆nick
func (r *TaobaosellercentersubuserspageAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosellercentersubuserspageAPIRequest) GetNick() string {
	return r._nick
}

// SetAccountStatus is AccountStatus Setter
// 可以不传，不传的话和老接口返回数据一样。如果传则根据子账号当前状态筛选 1正常   2冻结  3处罚，如果不传默认都返回
func (r *TaobaosellercentersubuserspageAPIRequest) SetAccountStatus(_accountStatus int64) error {
	r._accountStatus = _accountStatus
	r.Set("account_status", _accountStatus)
	return nil
}

// GetAccountStatus AccountStatus Getter
func (r TaobaosellercentersubuserspageAPIRequest) GetAccountStatus() int64 {
	return r._accountStatus
}

// SetPageSize is PageSize Setter
// 页大小，大于1小于200
func (r *TaobaosellercentersubuserspageAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosellercentersubuserspageAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 页码，大于等于1
func (r *TaobaosellercentersubuserspageAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaosellercentersubuserspageAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
