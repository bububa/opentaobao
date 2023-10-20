package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelbaseinfocitiesgetAPIRequest 【API3.0】度假线路商品发布时基础信息获取接口：地址数据查询 API请求
// taobao.alitrip.travel.baseinfo.cities.get
//
// 旅行度假新商品发布时可用的扩展接口，用于获取可用的出发地或目的地城市列表。
type TaobaoalitriptravelbaseinfocitiesgetAPIRequest struct {
	model.Params
	// 1-获取目的地城市列表 2-获取出发地城市列表
	_iocType int64
	// 1-境内跟团游 2-境内自由行 3-出境跟团游 4-出境自由行 5-境外当地玩乐 6-国际邮轮 9-境内邮轮
	_catType int64
}

// NewTaobaoalitriptravelbaseinfocitiesgetRequest 初始化TaobaoalitriptravelbaseinfocitiesgetAPIRequest对象
func NewTaobaoalitriptravelbaseinfocitiesgetRequest() *TaobaoalitriptravelbaseinfocitiesgetAPIRequest {
	return &TaobaoalitriptravelbaseinfocitiesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelbaseinfocitiesgetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.baseinfo.cities.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelbaseinfocitiesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelbaseinfocitiesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIocType is IocType Setter
// 1-获取目的地城市列表 2-获取出发地城市列表
func (r *TaobaoalitriptravelbaseinfocitiesgetAPIRequest) SetIocType(_iocType int64) error {
	r._iocType = _iocType
	r.Set("ioc_type", _iocType)
	return nil
}

// GetIocType IocType Getter
func (r TaobaoalitriptravelbaseinfocitiesgetAPIRequest) GetIocType() int64 {
	return r._iocType
}

// SetCatType is CatType Setter
// 1-境内跟团游 2-境内自由行 3-出境跟团游 4-出境自由行 5-境外当地玩乐 6-国际邮轮 9-境内邮轮
func (r *TaobaoalitriptravelbaseinfocitiesgetAPIRequest) SetCatType(_catType int64) error {
	r._catType = _catType
	r.Set("cat_type", _catType)
	return nil
}

// GetCatType CatType Getter
func (r TaobaoalitriptravelbaseinfocitiesgetAPIRequest) GetCatType() int64 {
	return r._catType
}
