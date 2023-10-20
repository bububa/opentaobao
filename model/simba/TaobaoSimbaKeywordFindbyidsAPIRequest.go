package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordFindbyidsAPIRequest （新）根据一堆关键词ids获取关键词 API请求
// taobao.simba.keyword.findbyids
//
// 根据一个关键词Id列表取得一组关键词
type TaobaoSimbaKeywordFindbyidsAPIRequest struct {
	model.Params
	// 关键词ids
	_bidwordIds []string
}

// NewTaobaoSimbaKeywordFindbyidsRequest 初始化TaobaoSimbaKeywordFindbyidsAPIRequest对象
func NewTaobaoSimbaKeywordFindbyidsRequest() *TaobaoSimbaKeywordFindbyidsAPIRequest {
	return &TaobaoSimbaKeywordFindbyidsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaKeywordFindbyidsAPIRequest) Reset() {
	r._bidwordIds = r._bidwordIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordFindbyidsAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keyword.findbyids"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordFindbyidsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaKeywordFindbyidsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordIds is BidwordIds Setter
// 关键词ids
func (r *TaobaoSimbaKeywordFindbyidsAPIRequest) SetBidwordIds(_bidwordIds []string) error {
	r._bidwordIds = _bidwordIds
	r.Set("bidword_ids", _bidwordIds)
	return nil
}

// GetBidwordIds BidwordIds Getter
func (r TaobaoSimbaKeywordFindbyidsAPIRequest) GetBidwordIds() []string {
	return r._bidwordIds
}

var poolTaobaoSimbaKeywordFindbyidsAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaKeywordFindbyidsRequest()
	},
}

// GetTaobaoSimbaKeywordFindbyidsRequest 从 sync.Pool 获取 TaobaoSimbaKeywordFindbyidsAPIRequest
func GetTaobaoSimbaKeywordFindbyidsAPIRequest() *TaobaoSimbaKeywordFindbyidsAPIRequest {
	return poolTaobaoSimbaKeywordFindbyidsAPIRequest.Get().(*TaobaoSimbaKeywordFindbyidsAPIRequest)
}

// ReleaseTaobaoSimbaKeywordFindbyidsAPIRequest 将 TaobaoSimbaKeywordFindbyidsAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaKeywordFindbyidsAPIRequest(v *TaobaoSimbaKeywordFindbyidsAPIRequest) {
	v.Reset()
	poolTaobaoSimbaKeywordFindbyidsAPIRequest.Put(v)
}
