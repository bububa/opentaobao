package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordupdateAPIRequest （新）关键词更新相关接口 API请求
// taobao.simba.keyword.update
//
// （新）关键词更新相关接口
type TaobaosimbakeywordupdateAPIRequest struct {
	model.Params
	// 关键词相关信息
	_bidwords []SiriusBidwordDto
}

// NewTaobaosimbakeywordupdateRequest 初始化TaobaosimbakeywordupdateAPIRequest对象
func NewTaobaosimbakeywordupdateRequest() *TaobaosimbakeywordupdateAPIRequest {
	return &TaobaosimbakeywordupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordupdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keyword.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwords is Bidwords Setter
// 关键词相关信息
func (r *TaobaosimbakeywordupdateAPIRequest) SetBidwords(_bidwords []SiriusBidwordDto) error {
	r._bidwords = _bidwords
	r.Set("bidwords", _bidwords)
	return nil
}

// GetBidwords Bidwords Getter
func (r TaobaosimbakeywordupdateAPIRequest) GetBidwords() []SiriusBidwordDto {
	return r._bidwords
}
