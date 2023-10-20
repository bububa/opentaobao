package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelcitycoordinatesbatchuploadAPIRequest 上传信息计算飞猪国际城市 API请求
// taobao.xhotel.city.coordinates.batch.upload
//
// 给供应商提供计算对应飞猪城市的服务，免去城市名称匹配流程，加快对接流程。目前只适用于国际城市，国内+港澳台暂不支持。
// 非实时计算接口，每次批量上传不少于1条的数据，后端离线计算，请于30分钟后再下载结果。
type TaobaoxhotelcitycoordinatesbatchuploadAPIRequest struct {
	model.Params
	// 经纬度列表
	_coordinateList []Coordinate
}

// NewTaobaoxhotelcitycoordinatesbatchuploadRequest 初始化TaobaoxhotelcitycoordinatesbatchuploadAPIRequest对象
func NewTaobaoxhotelcitycoordinatesbatchuploadRequest() *TaobaoxhotelcitycoordinatesbatchuploadAPIRequest {
	return &TaobaoxhotelcitycoordinatesbatchuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelcitycoordinatesbatchuploadAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.city.coordinates.batch.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelcitycoordinatesbatchuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelcitycoordinatesbatchuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCoordinateList is CoordinateList Setter
// 经纬度列表
func (r *TaobaoxhotelcitycoordinatesbatchuploadAPIRequest) SetCoordinateList(_coordinateList []Coordinate) error {
	r._coordinateList = _coordinateList
	r.Set("coordinate_list", _coordinateList)
	return nil
}

// GetCoordinateList CoordinateList Getter
func (r TaobaoxhotelcitycoordinatesbatchuploadAPIRequest) GetCoordinateList() []Coordinate {
	return r._coordinateList
}
