package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercenterSubusersPageAPIRequest 通过主账号登陆态分页查询子账号列表 API请求
// taobao.sellercenter.subusers.page
//
// 通过主账号登陆态分页查询子账号列表
type TaobaoSellercenterSubusersPageAPIRequest struct {
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

// NewTaobaoSellercenterSubusersPageRequest 初始化TaobaoSellercenterSubusersPageAPIRequest对象
func NewTaobaoSellercenterSubusersPageRequest() *TaobaoSellercenterSubusersPageAPIRequest {
	return &TaobaoSellercenterSubusersPageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSellercenterSubusersPageAPIRequest) GetApiMethodName() string {
	return "taobao.sellercenter.subusers.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSellercenterSubusersPageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSellercenterSubusersPageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主账号登陆nick
func (r *TaobaoSellercenterSubusersPageAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSellercenterSubusersPageAPIRequest) GetNick() string {
	return r._nick
}

// SetAccountStatus is AccountStatus Setter
// 可以不传，不传的话和老接口返回数据一样。如果传则根据子账号当前状态筛选 1正常   2冻结  3处罚，如果不传默认都返回
func (r *TaobaoSellercenterSubusersPageAPIRequest) SetAccountStatus(_accountStatus int64) error {
	r._accountStatus = _accountStatus
	r.Set("account_status", _accountStatus)
	return nil
}

// GetAccountStatus AccountStatus Getter
func (r TaobaoSellercenterSubusersPageAPIRequest) GetAccountStatus() int64 {
	return r._accountStatus
}

// SetPageSize is PageSize Setter
// 页大小，大于1小于200
func (r *TaobaoSellercenterSubusersPageAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSellercenterSubusersPageAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 页码，大于等于1
func (r *TaobaoSellercenterSubusersPageAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoSellercenterSubusersPageAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
