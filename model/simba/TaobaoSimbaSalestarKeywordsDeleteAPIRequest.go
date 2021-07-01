package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSalestarKeywordsDeleteAPIRequest
销量明星关键词删除 API请求
taobao.simba.salestar.keywords.delete

（新）关键词删除相关接口 */
type TaobaoSimbaSalestarKeywordsDeleteAPIRequest struct {
	model.Params
	// 关键词ids
	_bidwordIds []int64
}

// NewTaobaoSimbaSalestarKeywordsDeleteRequest 初始化TaobaoSimbaSalestarKeywordsDeleteAPIRequest对象
func NewTaobaoSimbaSalestarKeywordsDeleteRequest() *TaobaoSimbaSalestarKeywordsDeleteAPIRequest {
	return &TaobaoSimbaSalestarKeywordsDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarKeywordsDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.keywords.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarKeywordsDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BidwordIds Setter
// 关键词ids
func (r *TaobaoSimbaSalestarKeywordsDeleteAPIRequest) SetBidwordIds(_bidwordIds []int64) error {
	r._bidwordIds = _bidwordIds
	r.Set("bidword_ids", _bidwordIds)
	return nil
}

// Get BidwordIds Getter
func (r TaobaoSimbaSalestarKeywordsDeleteAPIRequest) GetBidwordIds() []int64 {
	return r._bidwordIds
}
