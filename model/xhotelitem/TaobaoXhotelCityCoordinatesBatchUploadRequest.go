package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传信息计算飞猪国际城市 API请求
taobao.xhotel.city.coordinates.batch.upload

给供应商提供计算对应飞猪城市的服务，免去城市名称匹配流程，加快对接流程。目前只适用于国际城市，国内+港澳台暂不支持。
非实时计算接口，每次批量上传不少于1条的数据，后端离线计算，请于30分钟后再下载结果。
*/
type TaobaoXhotelCityCoordinatesBatchUploadRequest struct {
    model.Params
    // 经纬度列表
    coordinateList   []Coordinate
}

// 初始化TaobaoXhotelCityCoordinatesBatchUploadRequest对象
func NewTaobaoXhotelCityCoordinatesBatchUploadRequest() *TaobaoXhotelCityCoordinatesBatchUploadRequest{
    return &TaobaoXhotelCityCoordinatesBatchUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCityCoordinatesBatchUploadRequest) GetApiMethodName() string {
    return "taobao.xhotel.city.coordinates.batch.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCityCoordinatesBatchUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CoordinateList Setter
// 经纬度列表
func (r *TaobaoXhotelCityCoordinatesBatchUploadRequest) SetCoordinateList(coordinateList []Coordinate) error {
    r.coordinateList = coordinateList
    r.Set("coordinate_list", coordinateList)
    return nil
}

// CoordinateList Getter
func (r TaobaoXhotelCityCoordinatesBatchUploadRequest) GetCoordinateList() []Coordinate {
    return r.coordinateList
}
