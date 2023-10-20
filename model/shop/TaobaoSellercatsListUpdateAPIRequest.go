package shop

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercatsListUpdateAPIRequest 更新卖家自定义类目 API请求
// taobao.sellercats.list.update
//
// 此API更新卖家店铺内自定义类目 &lt;br/&gt;注：因为缓存的关系，添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布
type TaobaoSellercatsListUpdateAPIRequest struct {
	model.Params
	// 卖家自定义类目名称。不超过20个字符
	_name string
	// 链接图片URL地址
	_pictUrl string
	// 卖家自定义类目编号
	_cid int64
	// 该类目在页面上的排序位置,取值范围:大于零的整数
	_sortOrder int64
}

// NewTaobaoSellercatsListUpdateRequest 初始化TaobaoSellercatsListUpdateAPIRequest对象
func NewTaobaoSellercatsListUpdateRequest() *TaobaoSellercatsListUpdateAPIRequest {
	return &TaobaoSellercatsListUpdateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSellercatsListUpdateAPIRequest) Reset() {
	r._name = ""
	r._pictUrl = ""
	r._cid = 0
	r._sortOrder = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSellercatsListUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.sellercats.list.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSellercatsListUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSellercatsListUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 卖家自定义类目名称。不超过20个字符
func (r *TaobaoSellercatsListUpdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoSellercatsListUpdateAPIRequest) GetName() string {
	return r._name
}

// SetPictUrl is PictUrl Setter
// 链接图片URL地址
func (r *TaobaoSellercatsListUpdateAPIRequest) SetPictUrl(_pictUrl string) error {
	r._pictUrl = _pictUrl
	r.Set("pict_url", _pictUrl)
	return nil
}

// GetPictUrl PictUrl Getter
func (r TaobaoSellercatsListUpdateAPIRequest) GetPictUrl() string {
	return r._pictUrl
}

// SetCid is Cid Setter
// 卖家自定义类目编号
func (r *TaobaoSellercatsListUpdateAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TaobaoSellercatsListUpdateAPIRequest) GetCid() int64 {
	return r._cid
}

// SetSortOrder is SortOrder Setter
// 该类目在页面上的排序位置,取值范围:大于零的整数
func (r *TaobaoSellercatsListUpdateAPIRequest) SetSortOrder(_sortOrder int64) error {
	r._sortOrder = _sortOrder
	r.Set("sort_order", _sortOrder)
	return nil
}

// GetSortOrder SortOrder Getter
func (r TaobaoSellercatsListUpdateAPIRequest) GetSortOrder() int64 {
	return r._sortOrder
}

var poolTaobaoSellercatsListUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSellercatsListUpdateRequest()
	},
}

// GetTaobaoSellercatsListUpdateRequest 从 sync.Pool 获取 TaobaoSellercatsListUpdateAPIRequest
func GetTaobaoSellercatsListUpdateAPIRequest() *TaobaoSellercatsListUpdateAPIRequest {
	return poolTaobaoSellercatsListUpdateAPIRequest.Get().(*TaobaoSellercatsListUpdateAPIRequest)
}

// ReleaseTaobaoSellercatsListUpdateAPIRequest 将 TaobaoSellercatsListUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSellercatsListUpdateAPIRequest(v *TaobaoSellercatsListUpdateAPIRequest) {
	v.Reset()
	poolTaobaoSellercatsListUpdateAPIRequest.Put(v)
}
