package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsBotsSkilsListAPIRequest 对外设备获取技能列表 API请求
// alibaba.ailabs.bots.skils.list
//
// 获取ai开放平台技能列表
type AlibabaAilabsBotsSkilsListAPIRequest struct {
	model.Params
	// 当前页
	_pageIndex int64
	// 分页单位
	_pageSize int64
}

// NewAlibabaAilabsBotsSkilsListRequest 初始化AlibabaAilabsBotsSkilsListAPIRequest对象
func NewAlibabaAilabsBotsSkilsListRequest() *AlibabaAilabsBotsSkilsListAPIRequest {
	return &AlibabaAilabsBotsSkilsListAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsBotsSkilsListAPIRequest) Reset() {
	r._pageIndex = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsBotsSkilsListAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.bots.skils.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsBotsSkilsListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsBotsSkilsListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageIndex is PageIndex Setter
// 当前页
func (r *AlibabaAilabsBotsSkilsListAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r AlibabaAilabsBotsSkilsListAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 分页单位
func (r *AlibabaAilabsBotsSkilsListAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAilabsBotsSkilsListAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaAilabsBotsSkilsListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsBotsSkilsListRequest()
	},
}

// GetAlibabaAilabsBotsSkilsListRequest 从 sync.Pool 获取 AlibabaAilabsBotsSkilsListAPIRequest
func GetAlibabaAilabsBotsSkilsListAPIRequest() *AlibabaAilabsBotsSkilsListAPIRequest {
	return poolAlibabaAilabsBotsSkilsListAPIRequest.Get().(*AlibabaAilabsBotsSkilsListAPIRequest)
}

// ReleaseAlibabaAilabsBotsSkilsListAPIRequest 将 AlibabaAilabsBotsSkilsListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsBotsSkilsListAPIRequest(v *AlibabaAilabsBotsSkilsListAPIRequest) {
	v.Reset()
	poolAlibabaAilabsBotsSkilsListAPIRequest.Put(v)
}
