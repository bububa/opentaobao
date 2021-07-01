package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaInsightCatsforecastnewGetAPIRequest
获取词的相关类目预测数据 API请求
taobao.simba.insight.catsforecastnew.get

根据给定的词，预测这些词的相关类目 */
type TaobaoSimbaInsightCatsforecastnewGetAPIRequest struct {
	model.Params
	// 需要查询的词列表
	_bidwordList []string
}

// NewTaobaoSimbaInsightCatsforecastnewGetRequest 初始化TaobaoSimbaInsightCatsforecastnewGetAPIRequest对象
func NewTaobaoSimbaInsightCatsforecastnewGetRequest() *TaobaoSimbaInsightCatsforecastnewGetAPIRequest {
	return &TaobaoSimbaInsightCatsforecastnewGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightCatsforecastnewGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.catsforecastnew.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightCatsforecastnewGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BidwordList Setter
// 需要查询的词列表
func (r *TaobaoSimbaInsightCatsforecastnewGetAPIRequest) SetBidwordList(_bidwordList []string) error {
	r._bidwordList = _bidwordList
	r.Set("bidword_list", _bidwordList)
	return nil
}

// Get BidwordList Getter
func (r TaobaoSimbaInsightCatsforecastnewGetAPIRequest) GetBidwordList() []string {
	return r._bidwordList
}
