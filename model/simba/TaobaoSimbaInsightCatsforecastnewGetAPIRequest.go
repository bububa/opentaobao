package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbainsightcatsforecastnewgetAPIRequest 获取词的相关类目预测数据 API请求
// taobao.simba.insight.catsforecastnew.get
//
// 根据给定的词，预测这些词的相关类目
type TaobaosimbainsightcatsforecastnewgetAPIRequest struct {
	model.Params
	// 需要查询的词列表
	_bidwordList []string
}

// NewTaobaosimbainsightcatsforecastnewgetRequest 初始化TaobaosimbainsightcatsforecastnewgetAPIRequest对象
func NewTaobaosimbainsightcatsforecastnewgetRequest() *TaobaosimbainsightcatsforecastnewgetAPIRequest {
	return &TaobaosimbainsightcatsforecastnewgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbainsightcatsforecastnewgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.catsforecastnew.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbainsightcatsforecastnewgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbainsightcatsforecastnewgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordList is BidwordList Setter
// 需要查询的词列表
func (r *TaobaosimbainsightcatsforecastnewgetAPIRequest) SetBidwordList(_bidwordList []string) error {
	r._bidwordList = _bidwordList
	r.Set("bidword_list", _bidwordList)
	return nil
}

// GetBidwordList BidwordList Getter
func (r TaobaosimbainsightcatsforecastnewgetAPIRequest) GetBidwordList() []string {
	return r._bidwordList
}
