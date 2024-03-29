package kbalgo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaKbalgoAlscpoisGetAPIRequest 百度批量获取本地poi接口 API请求
// alibaba.kbalgo.alscpois.get
//
// 接口用于百度方获取本地生活poi数据，分页获取。
type AlibabaKbalgoAlscpoisGetAPIRequest struct {
	model.Params
	// 第几页
	_pageNum int64
	// 每页的数量。
	_pageSize int64
}

// NewAlibabaKbalgoAlscpoisGetRequest 初始化AlibabaKbalgoAlscpoisGetAPIRequest对象
func NewAlibabaKbalgoAlscpoisGetRequest() *AlibabaKbalgoAlscpoisGetAPIRequest {
	return &AlibabaKbalgoAlscpoisGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaKbalgoAlscpoisGetAPIRequest) Reset() {
	r._pageNum = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaKbalgoAlscpoisGetAPIRequest) GetApiMethodName() string {
	return "alibaba.kbalgo.alscpois.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaKbalgoAlscpoisGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaKbalgoAlscpoisGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageNum is PageNum Setter
// 第几页
func (r *AlibabaKbalgoAlscpoisGetAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaKbalgoAlscpoisGetAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 每页的数量。
func (r *AlibabaKbalgoAlscpoisGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaKbalgoAlscpoisGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaKbalgoAlscpoisGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaKbalgoAlscpoisGetRequest()
	},
}

// GetAlibabaKbalgoAlscpoisGetRequest 从 sync.Pool 获取 AlibabaKbalgoAlscpoisGetAPIRequest
func GetAlibabaKbalgoAlscpoisGetAPIRequest() *AlibabaKbalgoAlscpoisGetAPIRequest {
	return poolAlibabaKbalgoAlscpoisGetAPIRequest.Get().(*AlibabaKbalgoAlscpoisGetAPIRequest)
}

// ReleaseAlibabaKbalgoAlscpoisGetAPIRequest 将 AlibabaKbalgoAlscpoisGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaKbalgoAlscpoisGetAPIRequest(v *AlibabaKbalgoAlscpoisGetAPIRequest) {
	v.Reset()
	poolAlibabaKbalgoAlscpoisGetAPIRequest.Put(v)
}
