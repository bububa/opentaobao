package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoProductsGetAPIRequest 获取产品列表 API请求
// taobao.products.get
//
// 根据淘宝会员帐号搜索所有产品信息，推荐使用taobao.products.search
// 注意：支持分页，每页最多返回100条,默认值为40,页码从1开始，默认为第一页
type TaobaoProductsGetAPIRequest struct {
	model.Params
	// 需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用","分隔
	_fields []string
	// 用户昵称
	_nick string
	// 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始.
	_pageNo int64
	// 每页条数.每页返回最多返回100条,默认值为40
	_pageSize int64
}

// NewTaobaoProductsGetRequest 初始化TaobaoProductsGetAPIRequest对象
func NewTaobaoProductsGetRequest() *TaobaoProductsGetAPIRequest {
	return &TaobaoProductsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoProductsGetAPIRequest) GetApiMethodName() string {
	return "taobao.products.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoProductsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoProductsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用&#34;,&#34;分隔
func (r *TaobaoProductsGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoProductsGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetNick is Nick Setter
// 用户昵称
func (r *TaobaoProductsGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoProductsGetAPIRequest) GetNick() string {
	return r._nick
}

// SetPageNo is PageNo Setter
// 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始.
func (r *TaobaoProductsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoProductsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数.每页返回最多返回100条,默认值为40
func (r *TaobaoProductsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoProductsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
