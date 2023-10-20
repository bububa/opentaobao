package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemQualificationDisplayGetAPIRequest 资质采集配置异步获取接口 API请求
// taobao.item.qualification.display.get
//
// 根据类目，商品，属性等参与动态获得资质采集配置
type TaobaoItemQualificationDisplayGetAPIRequest struct {
	model.Params
	// 参数列表，为key和value都是string的map的转化的json格式
	_param string
	// 商品id
	_itemId int64
	// 类目id
	_categoryId int64
}

// NewTaobaoItemQualificationDisplayGetRequest 初始化TaobaoItemQualificationDisplayGetAPIRequest对象
func NewTaobaoItemQualificationDisplayGetRequest() *TaobaoItemQualificationDisplayGetAPIRequest {
	return &TaobaoItemQualificationDisplayGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoItemQualificationDisplayGetAPIRequest) Reset() {
	r._param = ""
	r._itemId = 0
	r._categoryId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemQualificationDisplayGetAPIRequest) GetApiMethodName() string {
	return "taobao.item.qualification.display.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemQualificationDisplayGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemQualificationDisplayGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 参数列表，为key和value都是string的map的转化的json格式
func (r *TaobaoItemQualificationDisplayGetAPIRequest) SetParam(_param string) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoItemQualificationDisplayGetAPIRequest) GetParam() string {
	return r._param
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoItemQualificationDisplayGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoItemQualificationDisplayGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCategoryId is CategoryId Setter
// 类目id
func (r *TaobaoItemQualificationDisplayGetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TaobaoItemQualificationDisplayGetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

var poolTaobaoItemQualificationDisplayGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoItemQualificationDisplayGetRequest()
	},
}

// GetTaobaoItemQualificationDisplayGetRequest 从 sync.Pool 获取 TaobaoItemQualificationDisplayGetAPIRequest
func GetTaobaoItemQualificationDisplayGetAPIRequest() *TaobaoItemQualificationDisplayGetAPIRequest {
	return poolTaobaoItemQualificationDisplayGetAPIRequest.Get().(*TaobaoItemQualificationDisplayGetAPIRequest)
}

// ReleaseTaobaoItemQualificationDisplayGetAPIRequest 将 TaobaoItemQualificationDisplayGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoItemQualificationDisplayGetAPIRequest(v *TaobaoItemQualificationDisplayGetAPIRequest) {
	v.Reset()
	poolTaobaoItemQualificationDisplayGetAPIRequest.Put(v)
}
