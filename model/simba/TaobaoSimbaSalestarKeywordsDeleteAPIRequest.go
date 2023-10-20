package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasalestarkeywordsdeleteAPIRequest 销量明星关键词删除 API请求
// taobao.simba.salestar.keywords.delete
//
// （新）关键词删除相关接口
type TaobaosimbasalestarkeywordsdeleteAPIRequest struct {
	model.Params
	// 关键词ids
	_bidwordIds []string
}

// NewTaobaosimbasalestarkeywordsdeleteRequest 初始化TaobaosimbasalestarkeywordsdeleteAPIRequest对象
func NewTaobaosimbasalestarkeywordsdeleteRequest() *TaobaosimbasalestarkeywordsdeleteAPIRequest {
	return &TaobaosimbasalestarkeywordsdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbasalestarkeywordsdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.keywords.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbasalestarkeywordsdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbasalestarkeywordsdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordIds is BidwordIds Setter
// 关键词ids
func (r *TaobaosimbasalestarkeywordsdeleteAPIRequest) SetBidwordIds(_bidwordIds []string) error {
	r._bidwordIds = _bidwordIds
	r.Set("bidword_ids", _bidwordIds)
	return nil
}

// GetBidwordIds BidwordIds Getter
func (r TaobaosimbasalestarkeywordsdeleteAPIRequest) GetBidwordIds() []string {
	return r._bidwordIds
}
