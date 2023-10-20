package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaodistributoritemsgetAPIRequest 查询商品下载记录 API请求
// taobao.fenxiao.distributor.items.get
//
// 供应商查询分销商商品下载记录。
type TaobaofenxiaodistributoritemsgetAPIRequest struct {
	model.Params
	// 设置开始时间。空为不设置。
	_startModified string
	// 设置结束时间,空为不设置。
	_endModified string
	// 分销商ID 。
	_distributorId int64
	// 产品ID
	_productId int64
	// 每页记录数（默认20，最大50）
	_pageSize int64
	// 页码（大于0的整数，默认1）
	_pageNo int64
}

// NewTaobaofenxiaodistributoritemsgetRequest 初始化TaobaofenxiaodistributoritemsgetAPIRequest对象
func NewTaobaofenxiaodistributoritemsgetRequest() *TaobaofenxiaodistributoritemsgetAPIRequest {
	return &TaobaofenxiaodistributoritemsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaodistributoritemsgetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.distributor.items.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaodistributoritemsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaodistributoritemsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartModified is StartModified Setter
// 设置开始时间。空为不设置。
func (r *TaobaofenxiaodistributoritemsgetAPIRequest) SetStartModified(_startModified string) error {
	r._startModified = _startModified
	r.Set("start_modified", _startModified)
	return nil
}

// GetStartModified StartModified Getter
func (r TaobaofenxiaodistributoritemsgetAPIRequest) GetStartModified() string {
	return r._startModified
}

// SetEndModified is EndModified Setter
// 设置结束时间,空为不设置。
func (r *TaobaofenxiaodistributoritemsgetAPIRequest) SetEndModified(_endModified string) error {
	r._endModified = _endModified
	r.Set("end_modified", _endModified)
	return nil
}

// GetEndModified EndModified Getter
func (r TaobaofenxiaodistributoritemsgetAPIRequest) GetEndModified() string {
	return r._endModified
}

// SetDistributorId is DistributorId Setter
// 分销商ID 。
func (r *TaobaofenxiaodistributoritemsgetAPIRequest) SetDistributorId(_distributorId int64) error {
	r._distributorId = _distributorId
	r.Set("distributor_id", _distributorId)
	return nil
}

// GetDistributorId DistributorId Getter
func (r TaobaofenxiaodistributoritemsgetAPIRequest) GetDistributorId() int64 {
	return r._distributorId
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaofenxiaodistributoritemsgetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaofenxiaodistributoritemsgetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetPageSize is PageSize Setter
// 每页记录数（默认20，最大50）
func (r *TaobaofenxiaodistributoritemsgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaofenxiaodistributoritemsgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码（大于0的整数，默认1）
func (r *TaobaofenxiaodistributoritemsgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaofenxiaodistributoritemsgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
