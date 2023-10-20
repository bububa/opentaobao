package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubusersSubaccountSearchAPIRequest 根据子账号登录名后缀模糊搜索子账号列表 API请求
// taobao.subusers.subaccount.search
//
// 根据子账号冒号后缀搜索子账号列表，支持中文单字、英文单词（不支持英文单字母） 分词规则搜索，该搜索词必传。模糊搜索使用阿里云搜索引擎所以该接口增值收费，如果不需要模糊搜索仅需要分页获取子账号列表，请使用taobao.sellercenter.subusers.page接口
type TaobaoSubusersSubaccountSearchAPIRequest struct {
	model.Params
	// 主账号登录名
	_mainNick string
	// 根据子账号冒号后缀的搜索词，支持中文单字、英文单词 分词规则搜索。该搜索词必传。如果不需要模糊搜索仅需要分页获取子账号列表，请使用taobao.sellercenter.subusers.page接口
	_filterKey string
	// 页大小，大于1小于200
	_pageSize int64
	// 页码，大于等于1
	_pageNum int64
}

// NewTaobaoSubusersSubaccountSearchRequest 初始化TaobaoSubusersSubaccountSearchAPIRequest对象
func NewTaobaoSubusersSubaccountSearchRequest() *TaobaoSubusersSubaccountSearchAPIRequest {
	return &TaobaoSubusersSubaccountSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubusersSubaccountSearchAPIRequest) GetApiMethodName() string {
	return "taobao.subusers.subaccount.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubusersSubaccountSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubusersSubaccountSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainNick is MainNick Setter
// 主账号登录名
func (r *TaobaoSubusersSubaccountSearchAPIRequest) SetMainNick(_mainNick string) error {
	r._mainNick = _mainNick
	r.Set("main_nick", _mainNick)
	return nil
}

// GetMainNick MainNick Getter
func (r TaobaoSubusersSubaccountSearchAPIRequest) GetMainNick() string {
	return r._mainNick
}

// SetFilterKey is FilterKey Setter
// 根据子账号冒号后缀的搜索词，支持中文单字、英文单词 分词规则搜索。该搜索词必传。如果不需要模糊搜索仅需要分页获取子账号列表，请使用taobao.sellercenter.subusers.page接口
func (r *TaobaoSubusersSubaccountSearchAPIRequest) SetFilterKey(_filterKey string) error {
	r._filterKey = _filterKey
	r.Set("filter_key", _filterKey)
	return nil
}

// GetFilterKey FilterKey Getter
func (r TaobaoSubusersSubaccountSearchAPIRequest) GetFilterKey() string {
	return r._filterKey
}

// SetPageSize is PageSize Setter
// 页大小，大于1小于200
func (r *TaobaoSubusersSubaccountSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSubusersSubaccountSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 页码，大于等于1
func (r *TaobaoSubusersSubaccountSearchAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoSubusersSubaccountSearchAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
