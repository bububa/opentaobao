package lstvending

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendingGoodsSaveAPIRequest 自动售卖机商品回传 API请求
// alibaba.lst.vending.goods.save
//
// 零售通自动售卖机商品数据回流。
type AlibabaLstVendingGoodsSaveAPIRequest struct {
	model.Params
	// 商品信息
	_goodsDTOList []VendingGoodsDto
}

// NewAlibabaLstVendingGoodsSaveRequest 初始化AlibabaLstVendingGoodsSaveAPIRequest对象
func NewAlibabaLstVendingGoodsSaveRequest() *AlibabaLstVendingGoodsSaveAPIRequest {
	return &AlibabaLstVendingGoodsSaveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstVendingGoodsSaveAPIRequest) Reset() {
	r._goodsDTOList = r._goodsDTOList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingGoodsSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.goods.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingGoodsSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstVendingGoodsSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsDTOList is GoodsDTOList Setter
// 商品信息
func (r *AlibabaLstVendingGoodsSaveAPIRequest) SetGoodsDTOList(_goodsDTOList []VendingGoodsDto) error {
	r._goodsDTOList = _goodsDTOList
	r.Set("goods_d_t_o_list", _goodsDTOList)
	return nil
}

// GetGoodsDTOList GoodsDTOList Getter
func (r AlibabaLstVendingGoodsSaveAPIRequest) GetGoodsDTOList() []VendingGoodsDto {
	return r._goodsDTOList
}

var poolAlibabaLstVendingGoodsSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstVendingGoodsSaveRequest()
	},
}

// GetAlibabaLstVendingGoodsSaveRequest 从 sync.Pool 获取 AlibabaLstVendingGoodsSaveAPIRequest
func GetAlibabaLstVendingGoodsSaveAPIRequest() *AlibabaLstVendingGoodsSaveAPIRequest {
	return poolAlibabaLstVendingGoodsSaveAPIRequest.Get().(*AlibabaLstVendingGoodsSaveAPIRequest)
}

// ReleaseAlibabaLstVendingGoodsSaveAPIRequest 将 AlibabaLstVendingGoodsSaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstVendingGoodsSaveAPIRequest(v *AlibabaLstVendingGoodsSaveAPIRequest) {
	v.Reset()
	poolAlibabaLstVendingGoodsSaveAPIRequest.Put(v)
}
