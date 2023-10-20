package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbainsightrelatedwordsgetAPIRequest 获取词的相关词 API请求
// taobao.simba.insight.relatedwords.get
//
// 获取给定词的若干相关词，返回结果中越相关的权重越大，排在越前面，根据number参数对返回结果进行截断。
type TaobaosimbainsightrelatedwordsgetAPIRequest struct {
	model.Params
	// 要查询的词列表
	_bidwordList []string
	// 表示返回数据的条数
	_number int64
}

// NewTaobaosimbainsightrelatedwordsgetRequest 初始化TaobaosimbainsightrelatedwordsgetAPIRequest对象
func NewTaobaosimbainsightrelatedwordsgetRequest() *TaobaosimbainsightrelatedwordsgetAPIRequest {
	return &TaobaosimbainsightrelatedwordsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbainsightrelatedwordsgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.relatedwords.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbainsightrelatedwordsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbainsightrelatedwordsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordList is BidwordList Setter
// 要查询的词列表
func (r *TaobaosimbainsightrelatedwordsgetAPIRequest) SetBidwordList(_bidwordList []string) error {
	r._bidwordList = _bidwordList
	r.Set("bidword_list", _bidwordList)
	return nil
}

// GetBidwordList BidwordList Getter
func (r TaobaosimbainsightrelatedwordsgetAPIRequest) GetBidwordList() []string {
	return r._bidwordList
}

// SetNumber is Number Setter
// 表示返回数据的条数
func (r *TaobaosimbainsightrelatedwordsgetAPIRequest) SetNumber(_number int64) error {
	r._number = _number
	r.Set("number", _number)
	return nil
}

// GetNumber Number Getter
func (r TaobaosimbainsightrelatedwordsgetAPIRequest) GetNumber() int64 {
	return r._number
}
