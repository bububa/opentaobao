package kbalgo

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaKbalgoAlscpoisGetAPIRequest) GetApiMethodName() string {
	return "alibaba.kbalgo.alscpois.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaKbalgoAlscpoisGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
