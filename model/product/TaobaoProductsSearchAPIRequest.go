package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoProductsSearchAPIRequest 搜索产品信息 API请求
// taobao.products.search
//
// 只有天猫商家发布商品时才需要用到，并非商品搜索api，当前暂不提供商品搜索api。&lt;br/&gt;二种方式搜索所有产品信息(二种至少传一种): &lt;br/&gt;
// 传入关键字q搜索&lt;br/&gt;
// 传入cid和props搜索&lt;br/&gt;
// 返回值支持:product_id,name,pic_path,cid,props,price,tsc&lt;br/&gt;
// 当用户指定了cid并且cid为垂直市场（3C电器城、鞋城）的类目id时，默认只返回小二确认&lt;br/&gt;的产品。如果用户没有指定cid，或cid为普通的类目，默认返回商家确认或小二确认的产&lt;br/&gt;品。如果用户自定了status字段，以指定的status类型为准。
// &lt;br/&gt;新增一：
//    传入suite_items_str 按规格搜索套装产品。
//    返回字段增加suite_items_str,is_suite_effecitve支持。
type TaobaoProductsSearchAPIRequest struct {
	model.Params
	// 需返回的字段列表.可选值:Product数据结构中的以下字段:product_id,name,pic_url,cid,props,price,tsc;多个字段之间用","分隔.新增字段status(product的当前状态)
	_fields []string
	// 搜索的关键词是用来搜索产品的title.　注:q,cid和props至少传入一个
	_q string
	// 属性,属性值的组合.格式:pid:vid;pid:vid;调用taobao.itemprops.get获取类目属性pid <br/>,再用taobao.itempropvalues.get取得vid.
	_props string
	// 想要获取的产品的状态列表，支持多个状态并列获取，多个状态之间用","分隔，最多同时指定5种状态。例如，只获取小二确认的spu传入"3",只要商家确认的传入"0"，既要小二确认又要商家确认的传入"0,3"。目前只支持者两种类型的状态搜索，输入其他状态无效。
	_status string
	// 用户自定义关键属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234”
	_customerProps string
	// 按关联产品规格specs搜索套装产品
	_suiteItemsStr string
	// 按条码搜索产品信息,多个逗号隔开，不支持条码为全零的方式
	_barcodeStr string
	// 市场ID，1为取C2C市场的产品信息， 2为取B2C市场的产品信息。  不填写此值则默认取C2C的产品信息。
	_marketId string
	// 商品类目ID.调用taobao.itemcats.get获取.
	_cid int64
	// 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始.
	_pageNo int64
	// 每页条数.每页返回最多返回100条,默认值为40.
	_pageSize int64
	// 传入值为：3表示3C表示3C垂直市场产品，4表示鞋城垂直市场产品，8表示网游垂直市场产品。一次只能指定一种垂直市场类型
	_verticalMarket int64
}

// NewTaobaoProductsSearchRequest 初始化TaobaoProductsSearchAPIRequest对象
func NewTaobaoProductsSearchRequest() *TaobaoProductsSearchAPIRequest {
	return &TaobaoProductsSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoProductsSearchAPIRequest) GetApiMethodName() string {
	return "taobao.products.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoProductsSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 需返回的字段列表.可选值:Product数据结构中的以下字段:product_id,name,pic_url,cid,props,price,tsc;多个字段之间用&#34;,&#34;分隔.新增字段status(product的当前状态)
func (r *TaobaoProductsSearchAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoProductsSearchAPIRequest) GetFields() []string {
	return r._fields
}

// SetQ is Q Setter
// 搜索的关键词是用来搜索产品的title.　注:q,cid和props至少传入一个
func (r *TaobaoProductsSearchAPIRequest) SetQ(_q string) error {
	r._q = _q
	r.Set("q", _q)
	return nil
}

// GetQ Q Getter
func (r TaobaoProductsSearchAPIRequest) GetQ() string {
	return r._q
}

// SetProps is Props Setter
// 属性,属性值的组合.格式:pid:vid;pid:vid;调用taobao.itemprops.get获取类目属性pid &lt;br/&gt;,再用taobao.itempropvalues.get取得vid.
func (r *TaobaoProductsSearchAPIRequest) SetProps(_props string) error {
	r._props = _props
	r.Set("props", _props)
	return nil
}

// GetProps Props Getter
func (r TaobaoProductsSearchAPIRequest) GetProps() string {
	return r._props
}

// SetStatus is Status Setter
// 想要获取的产品的状态列表，支持多个状态并列获取，多个状态之间用&#34;,&#34;分隔，最多同时指定5种状态。例如，只获取小二确认的spu传入&#34;3&#34;,只要商家确认的传入&#34;0&#34;，既要小二确认又要商家确认的传入&#34;0,3&#34;。目前只支持者两种类型的状态搜索，输入其他状态无效。
func (r *TaobaoProductsSearchAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoProductsSearchAPIRequest) GetStatus() string {
	return r._status
}

// SetCustomerProps is CustomerProps Setter
// 用户自定义关键属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234”
func (r *TaobaoProductsSearchAPIRequest) SetCustomerProps(_customerProps string) error {
	r._customerProps = _customerProps
	r.Set("customer_props", _customerProps)
	return nil
}

// GetCustomerProps CustomerProps Getter
func (r TaobaoProductsSearchAPIRequest) GetCustomerProps() string {
	return r._customerProps
}

// SetSuiteItemsStr is SuiteItemsStr Setter
// 按关联产品规格specs搜索套装产品
func (r *TaobaoProductsSearchAPIRequest) SetSuiteItemsStr(_suiteItemsStr string) error {
	r._suiteItemsStr = _suiteItemsStr
	r.Set("suite_items_str", _suiteItemsStr)
	return nil
}

// GetSuiteItemsStr SuiteItemsStr Getter
func (r TaobaoProductsSearchAPIRequest) GetSuiteItemsStr() string {
	return r._suiteItemsStr
}

// SetBarcodeStr is BarcodeStr Setter
// 按条码搜索产品信息,多个逗号隔开，不支持条码为全零的方式
func (r *TaobaoProductsSearchAPIRequest) SetBarcodeStr(_barcodeStr string) error {
	r._barcodeStr = _barcodeStr
	r.Set("barcode_str", _barcodeStr)
	return nil
}

// GetBarcodeStr BarcodeStr Getter
func (r TaobaoProductsSearchAPIRequest) GetBarcodeStr() string {
	return r._barcodeStr
}

// SetMarketId is MarketId Setter
// 市场ID，1为取C2C市场的产品信息， 2为取B2C市场的产品信息。  不填写此值则默认取C2C的产品信息。
func (r *TaobaoProductsSearchAPIRequest) SetMarketId(_marketId string) error {
	r._marketId = _marketId
	r.Set("market_id", _marketId)
	return nil
}

// GetMarketId MarketId Getter
func (r TaobaoProductsSearchAPIRequest) GetMarketId() string {
	return r._marketId
}

// SetCid is Cid Setter
// 商品类目ID.调用taobao.itemcats.get获取.
func (r *TaobaoProductsSearchAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TaobaoProductsSearchAPIRequest) GetCid() int64 {
	return r._cid
}

// SetPageNo is PageNo Setter
// 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始.
func (r *TaobaoProductsSearchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoProductsSearchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数.每页返回最多返回100条,默认值为40.
func (r *TaobaoProductsSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoProductsSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetVerticalMarket is VerticalMarket Setter
// 传入值为：3表示3C表示3C垂直市场产品，4表示鞋城垂直市场产品，8表示网游垂直市场产品。一次只能指定一种垂直市场类型
func (r *TaobaoProductsSearchAPIRequest) SetVerticalMarket(_verticalMarket int64) error {
	r._verticalMarket = _verticalMarket
	r.Set("vertical_market", _verticalMarket)
	return nil
}

// GetVerticalMarket VerticalMarket Getter
func (r TaobaoProductsSearchAPIRequest) GetVerticalMarket() int64 {
	return r._verticalMarket
}
