package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest 单品推广搜索人群修改溢价 API请求
// taobao.simba.serchcrowd.price.batch.update
//
// 单品推广搜索人群修改溢价, 不支持跨推广单元修改
type TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest struct {
	model.Params
	// 被操作者的淘宝昵称
	_nick string
	// 子帐号nick
	_subNick string
	// 需要修改出价的人群包id,批量传入的时候用,分割
	_adgroupCrowdIds []int64
	// 推广单元id
	_adgroupId int64
	// 人群溢价比例，溢价范围[5,300]
	_discount int64
}

// NewTaobaoSimbaSerchcrowdPriceBatchUpdateRequest 初始化TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest对象
func NewTaobaoSimbaSerchcrowdPriceBatchUpdateRequest() *TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest {
	return &TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.serchcrowd.price.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest) GetNick() string {
	return r._nick
}

// Set is SubNick Setter
// 子帐号nick
func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest) SetSubNick(_subNick string) error {
	r._subNick = _subNick
	r.Set("sub_nick", _subNick)
	return nil
}

// Get SubNick Getter
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest) GetSubNick() string {
	return r._subNick
}

// Set is AdgroupCrowdIds Setter
// 需要修改出价的人群包id,批量传入的时候用,分割
func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest) SetAdgroupCrowdIds(_adgroupCrowdIds []int64) error {
	r._adgroupCrowdIds = _adgroupCrowdIds
	r.Set("adgroup_crowd_ids", _adgroupCrowdIds)
	return nil
}

// Get AdgroupCrowdIds Getter
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest) GetAdgroupCrowdIds() []int64 {
	return r._adgroupCrowdIds
}

// Set is AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// Set is Discount Setter
// 人群溢价比例，溢价范围[5,300]
func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest) SetDiscount(_discount int64) error {
	r._discount = _discount
	r.Set("discount", _discount)
	return nil
}

// Get Discount Getter
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest) GetDiscount() int64 {
	return r._discount
}
