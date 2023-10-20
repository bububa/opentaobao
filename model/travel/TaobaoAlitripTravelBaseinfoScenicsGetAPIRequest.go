package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest 【API3.0】基础信息获取接口：景点数据查询 API请求
// taobao.alitrip.travel.baseinfo.scenics.get
//
// 商品发布辅助接口，用于飞猪度假或门票商品发布时 获取可用的景点（及景点下收费项目）信息列表。
type TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest struct {
	model.Params
	// 城市名称
	_city string
	// 景点简称
	_scenic string
	// 景点id，可选。若传了该值，则查询结果中只会保留该id的景点，其余景点信息将被过滤
	_scenicId int64
}

// NewTaobaoAlitripTravelBaseinfoScenicsGetRequest 初始化TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest对象
func NewTaobaoAlitripTravelBaseinfoScenicsGetRequest() *TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest {
	return &TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) Reset() {
	r._city = ""
	r._scenic = ""
	r._scenicId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.baseinfo.scenics.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCity is City Setter
// 城市名称
func (r *TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) GetCity() string {
	return r._city
}

// SetScenic is Scenic Setter
// 景点简称
func (r *TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) SetScenic(_scenic string) error {
	r._scenic = _scenic
	r.Set("scenic", _scenic)
	return nil
}

// GetScenic Scenic Getter
func (r TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) GetScenic() string {
	return r._scenic
}

// SetScenicId is ScenicId Setter
// 景点id，可选。若传了该值，则查询结果中只会保留该id的景点，其余景点信息将被过滤
func (r *TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) SetScenicId(_scenicId int64) error {
	r._scenicId = _scenicId
	r.Set("scenic_id", _scenicId)
	return nil
}

// GetScenicId ScenicId Getter
func (r TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) GetScenicId() int64 {
	return r._scenicId
}

var poolTaobaoAlitripTravelBaseinfoScenicsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelBaseinfoScenicsGetRequest()
	},
}

// GetTaobaoAlitripTravelBaseinfoScenicsGetRequest 从 sync.Pool 获取 TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest
func GetTaobaoAlitripTravelBaseinfoScenicsGetAPIRequest() *TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest {
	return poolTaobaoAlitripTravelBaseinfoScenicsGetAPIRequest.Get().(*TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest)
}

// ReleaseTaobaoAlitripTravelBaseinfoScenicsGetAPIRequest 将 TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelBaseinfoScenicsGetAPIRequest(v *TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelBaseinfoScenicsGetAPIRequest.Put(v)
}
