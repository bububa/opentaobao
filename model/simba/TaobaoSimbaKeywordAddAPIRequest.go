package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordaddAPIRequest （新）关键词新增接口 API请求
// taobao.simba.keyword.add
//
// （新）关键词更新相关接口
type TaobaosimbakeywordaddAPIRequest struct {
	model.Params
	// 关键词相关信息
	_bidwords []SiriusBidwordDto
	// 推广单元id
	_adgroupId int64
}

// NewTaobaosimbakeywordaddRequest 初始化TaobaosimbakeywordaddAPIRequest对象
func NewTaobaosimbakeywordaddRequest() *TaobaosimbakeywordaddAPIRequest {
	return &TaobaosimbakeywordaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordaddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keyword.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwords is Bidwords Setter
// 关键词相关信息
func (r *TaobaosimbakeywordaddAPIRequest) SetBidwords(_bidwords []SiriusBidwordDto) error {
	r._bidwords = _bidwords
	r.Set("bidwords", _bidwords)
	return nil
}

// GetBidwords Bidwords Getter
func (r TaobaosimbakeywordaddAPIRequest) GetBidwords() []SiriusBidwordDto {
	return r._bidwords
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaosimbakeywordaddAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbakeywordaddAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
