package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaInsightCatsforecastnewGetAPIRequest 获取词的相关类目预测数据 API请求
// taobao.simba.insight.catsforecastnew.get
//
// 根据给定的词，预测这些词的相关类目
type TaobaoSimbaInsightCatsforecastnewGetAPIRequest struct {
	model.Params
	// 需要查询的词列表
	_bidwordList []string
}

// NewTaobaoSimbaInsightCatsforecastnewGetRequest 初始化TaobaoSimbaInsightCatsforecastnewGetAPIRequest对象
func NewTaobaoSimbaInsightCatsforecastnewGetRequest() *TaobaoSimbaInsightCatsforecastnewGetAPIRequest {
	return &TaobaoSimbaInsightCatsforecastnewGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaInsightCatsforecastnewGetAPIRequest) Reset() {
	r._bidwordList = r._bidwordList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightCatsforecastnewGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.catsforecastnew.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightCatsforecastnewGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaInsightCatsforecastnewGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordList is BidwordList Setter
// 需要查询的词列表
func (r *TaobaoSimbaInsightCatsforecastnewGetAPIRequest) SetBidwordList(_bidwordList []string) error {
	r._bidwordList = _bidwordList
	r.Set("bidword_list", _bidwordList)
	return nil
}

// GetBidwordList BidwordList Getter
func (r TaobaoSimbaInsightCatsforecastnewGetAPIRequest) GetBidwordList() []string {
	return r._bidwordList
}

var poolTaobaoSimbaInsightCatsforecastnewGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaInsightCatsforecastnewGetRequest()
	},
}

// GetTaobaoSimbaInsightCatsforecastnewGetRequest 从 sync.Pool 获取 TaobaoSimbaInsightCatsforecastnewGetAPIRequest
func GetTaobaoSimbaInsightCatsforecastnewGetAPIRequest() *TaobaoSimbaInsightCatsforecastnewGetAPIRequest {
	return poolTaobaoSimbaInsightCatsforecastnewGetAPIRequest.Get().(*TaobaoSimbaInsightCatsforecastnewGetAPIRequest)
}

// ReleaseTaobaoSimbaInsightCatsforecastnewGetAPIRequest 将 TaobaoSimbaInsightCatsforecastnewGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaInsightCatsforecastnewGetAPIRequest(v *TaobaoSimbaInsightCatsforecastnewGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaInsightCatsforecastnewGetAPIRequest.Put(v)
}
