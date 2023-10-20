package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoProductGetAPIRequest 获取一个产品的信息 API请求
// taobao.product.get
//
// 天猫商家发布商品时，查询关联产品信息时使用，非商品查询接口。商品查询接口：taobao.item.seller.get&lt;br&gt;
// 两种方式查看一个产品详细信息:
// 传入product_id来查询；传入cid和props来查询
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoProductGetAPIRequest struct {
	model.Params
	// 需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用","分隔.
	_fields string
	// 比如:诺基亚N73这个产品的关键属性列表就是:品牌:诺基亚;型号:N73,对应的PV值就是10005:10027;10006:29729.
	_props string
	// Product的id.两种方式来查看一个产品:1.传入product_id来查询 2.传入cid和props来查询
	_productId int64
	// 商品类目id.调用taobao.itemcats.get获取;必须是叶子类目id,如果没有传product_id,那么cid和props必须要传.
	_cid int64
}

// NewTaobaoProductGetRequest 初始化TaobaoProductGetAPIRequest对象
func NewTaobaoProductGetRequest() *TaobaoProductGetAPIRequest {
	return &TaobaoProductGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoProductGetAPIRequest) Reset() {
	r._fields = ""
	r._props = ""
	r._productId = 0
	r._cid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoProductGetAPIRequest) GetApiMethodName() string {
	return "taobao.product.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoProductGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoProductGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用&#34;,&#34;分隔.
func (r *TaobaoProductGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoProductGetAPIRequest) GetFields() string {
	return r._fields
}

// SetProps is Props Setter
// 比如:诺基亚N73这个产品的关键属性列表就是:品牌:诺基亚;型号:N73,对应的PV值就是10005:10027;10006:29729.
func (r *TaobaoProductGetAPIRequest) SetProps(_props string) error {
	r._props = _props
	r.Set("props", _props)
	return nil
}

// GetProps Props Getter
func (r TaobaoProductGetAPIRequest) GetProps() string {
	return r._props
}

// SetProductId is ProductId Setter
// Product的id.两种方式来查看一个产品:1.传入product_id来查询 2.传入cid和props来查询
func (r *TaobaoProductGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoProductGetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetCid is Cid Setter
// 商品类目id.调用taobao.itemcats.get获取;必须是叶子类目id,如果没有传product_id,那么cid和props必须要传.
func (r *TaobaoProductGetAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TaobaoProductGetAPIRequest) GetCid() int64 {
	return r._cid
}

var poolTaobaoProductGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoProductGetRequest()
	},
}

// GetTaobaoProductGetRequest 从 sync.Pool 获取 TaobaoProductGetAPIRequest
func GetTaobaoProductGetAPIRequest() *TaobaoProductGetAPIRequest {
	return poolTaobaoProductGetAPIRequest.Get().(*TaobaoProductGetAPIRequest)
}

// ReleaseTaobaoProductGetAPIRequest 将 TaobaoProductGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoProductGetAPIRequest(v *TaobaoProductGetAPIRequest) {
	v.Reset()
	poolTaobaoProductGetAPIRequest.Put(v)
}
