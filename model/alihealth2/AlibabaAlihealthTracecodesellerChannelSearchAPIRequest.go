package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellerchannelsearchAPIRequest 查询渠道商api API请求
// alibaba.alihealth.tracecodeseller.channel.search
//
// 查询渠道商api
type AlibabaalihealthtracecodesellerchannelsearchAPIRequest struct {
	model.Params
	// 身份认证
	_skeyCode string
	// 商家id
	_entInfoId int64
	// 0 出库 2 入库
	_outInType int64
	// 第几页
	_page int64
	// 每页几条
	_pageSize int64
}

// NewAlibabaalihealthtracecodesellerchannelsearchRequest 初始化AlibabaalihealthtracecodesellerchannelsearchAPIRequest对象
func NewAlibabaalihealthtracecodesellerchannelsearchRequest() *AlibabaalihealthtracecodesellerchannelsearchAPIRequest {
	return &AlibabaalihealthtracecodesellerchannelsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtracecodesellerchannelsearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.channel.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtracecodesellerchannelsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtracecodesellerchannelsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkeyCode is SkeyCode Setter
// 身份认证
func (r *AlibabaalihealthtracecodesellerchannelsearchAPIRequest) SetSkeyCode(_skeyCode string) error {
	r._skeyCode = _skeyCode
	r.Set("skey_code", _skeyCode)
	return nil
}

// GetSkeyCode SkeyCode Getter
func (r AlibabaalihealthtracecodesellerchannelsearchAPIRequest) GetSkeyCode() string {
	return r._skeyCode
}

// SetEntInfoId is EntInfoId Setter
// 商家id
func (r *AlibabaalihealthtracecodesellerchannelsearchAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// GetEntInfoId EntInfoId Getter
func (r AlibabaalihealthtracecodesellerchannelsearchAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}

// SetOutInType is OutInType Setter
// 0 出库 2 入库
func (r *AlibabaalihealthtracecodesellerchannelsearchAPIRequest) SetOutInType(_outInType int64) error {
	r._outInType = _outInType
	r.Set("out_in_type", _outInType)
	return nil
}

// GetOutInType OutInType Getter
func (r AlibabaalihealthtracecodesellerchannelsearchAPIRequest) GetOutInType() int64 {
	return r._outInType
}

// SetPage is Page Setter
// 第几页
func (r *AlibabaalihealthtracecodesellerchannelsearchAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaalihealthtracecodesellerchannelsearchAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 每页几条
func (r *AlibabaalihealthtracecodesellerchannelsearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaalihealthtracecodesellerchannelsearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
