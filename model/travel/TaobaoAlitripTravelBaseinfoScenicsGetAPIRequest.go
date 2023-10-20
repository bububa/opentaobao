package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelbaseinfoscenicsgetAPIRequest 【API3.0】基础信息获取接口：景点数据查询 API请求
// taobao.alitrip.travel.baseinfo.scenics.get
//
// 商品发布辅助接口，用于飞猪度假或门票商品发布时 获取可用的景点（及景点下收费项目）信息列表。
type TaobaoalitriptravelbaseinfoscenicsgetAPIRequest struct {
	model.Params
	// 城市名称
	_city string
	// 景点简称
	_scenic string
	// 景点id，可选。若传了该值，则查询结果中只会保留该id的景点，其余景点信息将被过滤
	_scenicId int64
}

// NewTaobaoalitriptravelbaseinfoscenicsgetRequest 初始化TaobaoalitriptravelbaseinfoscenicsgetAPIRequest对象
func NewTaobaoalitriptravelbaseinfoscenicsgetRequest() *TaobaoalitriptravelbaseinfoscenicsgetAPIRequest {
	return &TaobaoalitriptravelbaseinfoscenicsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelbaseinfoscenicsgetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.baseinfo.scenics.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelbaseinfoscenicsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelbaseinfoscenicsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCity is City Setter
// 城市名称
func (r *TaobaoalitriptravelbaseinfoscenicsgetAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r TaobaoalitriptravelbaseinfoscenicsgetAPIRequest) GetCity() string {
	return r._city
}

// SetScenic is Scenic Setter
// 景点简称
func (r *TaobaoalitriptravelbaseinfoscenicsgetAPIRequest) SetScenic(_scenic string) error {
	r._scenic = _scenic
	r.Set("scenic", _scenic)
	return nil
}

// GetScenic Scenic Getter
func (r TaobaoalitriptravelbaseinfoscenicsgetAPIRequest) GetScenic() string {
	return r._scenic
}

// SetScenicId is ScenicId Setter
// 景点id，可选。若传了该值，则查询结果中只会保留该id的景点，其余景点信息将被过滤
func (r *TaobaoalitriptravelbaseinfoscenicsgetAPIRequest) SetScenicId(_scenicId int64) error {
	r._scenicId = _scenicId
	r.Set("scenic_id", _scenicId)
	return nil
}

// GetScenicId ScenicId Getter
func (r TaobaoalitriptravelbaseinfoscenicsgetAPIRequest) GetScenicId() int64 {
	return r._scenicId
}
