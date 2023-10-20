package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaovmarketetickettasksgetAPIRequest 任务列表获取接口 API请求
// taobao.vmarket.eticket.tasks.get
//
// 外部合作卖家获取任务列表的信息：如发码同通知失败或者回调失败的订单号
type TaobaovmarketetickettasksgetAPIRequest struct {
	model.Params
	// 卖家家ID(信任卖家不必传，码商可选)
	_sellerId int64
	// 返回结果类型:<br/>1:返回通知失败的订单<br/>2.返回通知成功回调失败的订单
	_type int64
	// 页码。取值范围:大于零的整数; 默认值:1
	_pageNo int64
	// 每页获取条数。默认值40，最小值1，最大值100。
	_pageSize int64
	// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
	_codemerchantId int64
}

// NewTaobaovmarketetickettasksgetRequest 初始化TaobaovmarketetickettasksgetAPIRequest对象
func NewTaobaovmarketetickettasksgetRequest() *TaobaovmarketetickettasksgetAPIRequest {
	return &TaobaovmarketetickettasksgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaovmarketetickettasksgetAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.tasks.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaovmarketetickettasksgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaovmarketetickettasksgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerId is SellerId Setter
// 卖家家ID(信任卖家不必传，码商可选)
func (r *TaobaovmarketetickettasksgetAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TaobaovmarketetickettasksgetAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// SetType is Type Setter
// 返回结果类型:&lt;br/&gt;1:返回通知失败的订单&lt;br/&gt;2.返回通知成功回调失败的订单
func (r *TaobaovmarketetickettasksgetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaovmarketetickettasksgetAPIRequest) GetType() int64 {
	return r._type
}

// SetPageNo is PageNo Setter
// 页码。取值范围:大于零的整数; 默认值:1
func (r *TaobaovmarketetickettasksgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaovmarketetickettasksgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页获取条数。默认值40，最小值1，最大值100。
func (r *TaobaovmarketetickettasksgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaovmarketetickettasksgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCodemerchantId is CodemerchantId Setter
// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
func (r *TaobaovmarketetickettasksgetAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// GetCodemerchantId CodemerchantId Getter
func (r TaobaovmarketetickettasksgetAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}
