package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosellercatslistupdateAPIRequest 更新卖家自定义类目 API请求
// taobao.sellercats.list.update
//
// 此API更新卖家店铺内自定义类目 &lt;br/&gt;注：因为缓存的关系，添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布
type TaobaosellercatslistupdateAPIRequest struct {
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

// NewTaobaosellercatslistupdateRequest 初始化TaobaosellercatslistupdateAPIRequest对象
func NewTaobaosellercatslistupdateRequest() *TaobaosellercatslistupdateAPIRequest {
	return &TaobaosellercatslistupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosellercatslistupdateAPIRequest) GetApiMethodName() string {
	return "taobao.sellercats.list.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosellercatslistupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosellercatslistupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 卖家自定义类目名称。不超过20个字符
func (r *TaobaosellercatslistupdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaosellercatslistupdateAPIRequest) GetName() string {
	return r._name
}

// SetPictUrl is PictUrl Setter
// 链接图片URL地址
func (r *TaobaosellercatslistupdateAPIRequest) SetPictUrl(_pictUrl string) error {
	r._pictUrl = _pictUrl
	r.Set("pict_url", _pictUrl)
	return nil
}

// GetPictUrl PictUrl Getter
func (r TaobaosellercatslistupdateAPIRequest) GetPictUrl() string {
	return r._pictUrl
}

// SetCid is Cid Setter
// 卖家自定义类目编号
func (r *TaobaosellercatslistupdateAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TaobaosellercatslistupdateAPIRequest) GetCid() int64 {
	return r._cid
}

// SetSortOrder is SortOrder Setter
// 该类目在页面上的排序位置,取值范围:大于零的整数
func (r *TaobaosellercatslistupdateAPIRequest) SetSortOrder(_sortOrder int64) error {
	r._sortOrder = _sortOrder
	r.Set("sort_order", _sortOrder)
	return nil
}

// GetSortOrder SortOrder Getter
func (r TaobaosellercatslistupdateAPIRequest) GetSortOrder() int64 {
	return r._sortOrder
}
