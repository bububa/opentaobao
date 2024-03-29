package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceFloorGetbybuildingidAPIRequest 根据楼宇ID获取楼层 API请求
// alibaba.campus.space.floor.getbybuildingid
//
// 根据楼宇ID获取楼层
// HSF接口名称：com.alibaba.campus.api.space.service.top.FloorApiTopService
// HSF方法名称：getFloorList
type AlibabaCampusSpaceFloorGetbybuildingidAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 楼宇iD封装
	_param1 *FloorQuery
}

// NewAlibabaCampusSpaceFloorGetbybuildingidRequest 初始化AlibabaCampusSpaceFloorGetbybuildingidAPIRequest对象
func NewAlibabaCampusSpaceFloorGetbybuildingidRequest() *AlibabaCampusSpaceFloorGetbybuildingidAPIRequest {
	return &AlibabaCampusSpaceFloorGetbybuildingidAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.floor.getbybuildingid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 楼宇iD封装
func (r *AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) SetParam1(_param1 *FloorQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) GetParam1() *FloorQuery {
	return r._param1
}

var poolAlibabaCampusSpaceFloorGetbybuildingidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusSpaceFloorGetbybuildingidRequest()
	},
}

// GetAlibabaCampusSpaceFloorGetbybuildingidRequest 从 sync.Pool 获取 AlibabaCampusSpaceFloorGetbybuildingidAPIRequest
func GetAlibabaCampusSpaceFloorGetbybuildingidAPIRequest() *AlibabaCampusSpaceFloorGetbybuildingidAPIRequest {
	return poolAlibabaCampusSpaceFloorGetbybuildingidAPIRequest.Get().(*AlibabaCampusSpaceFloorGetbybuildingidAPIRequest)
}

// ReleaseAlibabaCampusSpaceFloorGetbybuildingidAPIRequest 将 AlibabaCampusSpaceFloorGetbybuildingidAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusSpaceFloorGetbybuildingidAPIRequest(v *AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) {
	v.Reset()
	poolAlibabaCampusSpaceFloorGetbybuildingidAPIRequest.Put(v)
}
