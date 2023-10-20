package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerChannelSearchAPIRequest 查询渠道商api API请求
// alibaba.alihealth.tracecodeseller.channel.search
//
// 查询渠道商api
type AlibabaAlihealthTracecodesellerChannelSearchAPIRequest struct {
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

// NewAlibabaAlihealthTracecodesellerChannelSearchRequest 初始化AlibabaAlihealthTracecodesellerChannelSearchAPIRequest对象
func NewAlibabaAlihealthTracecodesellerChannelSearchRequest() *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest {
	return &AlibabaAlihealthTracecodesellerChannelSearchAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) Reset() {
	r._skeyCode = ""
	r._entInfoId = 0
	r._outInType = 0
	r._page = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.channel.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkeyCode is SkeyCode Setter
// 身份认证
func (r *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) SetSkeyCode(_skeyCode string) error {
	r._skeyCode = _skeyCode
	r.Set("skey_code", _skeyCode)
	return nil
}

// GetSkeyCode SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetSkeyCode() string {
	return r._skeyCode
}

// SetEntInfoId is EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// GetEntInfoId EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}

// SetOutInType is OutInType Setter
// 0 出库 2 入库
func (r *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) SetOutInType(_outInType int64) error {
	r._outInType = _outInType
	r.Set("out_in_type", _outInType)
	return nil
}

// GetOutInType OutInType Getter
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetOutInType() int64 {
	return r._outInType
}

// SetPage is Page Setter
// 第几页
func (r *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 每页几条
func (r *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaAlihealthTracecodesellerChannelSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthTracecodesellerChannelSearchRequest()
	},
}

// GetAlibabaAlihealthTracecodesellerChannelSearchRequest 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerChannelSearchAPIRequest
func GetAlibabaAlihealthTracecodesellerChannelSearchAPIRequest() *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest {
	return poolAlibabaAlihealthTracecodesellerChannelSearchAPIRequest.Get().(*AlibabaAlihealthTracecodesellerChannelSearchAPIRequest)
}

// ReleaseAlibabaAlihealthTracecodesellerChannelSearchAPIRequest 将 AlibabaAlihealthTracecodesellerChannelSearchAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerChannelSearchAPIRequest(v *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerChannelSearchAPIRequest.Put(v)
}
