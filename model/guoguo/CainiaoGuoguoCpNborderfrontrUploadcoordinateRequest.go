package guoguo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传小件员GPS位置信息 API请求
cainiao.guoguo.cp.nborderfrontr.uploadcoordinate

上传小件员GPS位置信息
*/
type CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest struct {
    model.Params
    // 小件员所在公司编号
    _cpCode   string
    // 小件员员工编号
    _cpUserId   string
    // 上报时间，格式：yyyy-MM-dd HH:mm:ss
    _timeStamp   string
    // 来源：1.小件员app sdk 2.驿站 3. 裹裹 10001.圆通行者
    _source   int64
    // 经度
    _lng   string
    // 纬度
    _lat   string
    // 0 安卓定位，     1 苹果定位，  2 其他系统定位，   10 高德定位，  11 百度定位，  12 google定位     13 其他
    _gpsType   string
}

// 初始化CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest对象
func NewCainiaoGuoguoCpNborderfrontrUploadcoordinateRequest() *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest{
    return &CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetApiMethodName() string {
    return "cainiao.guoguo.cp.nborderfrontr.uploadcoordinate"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CpCode Setter
// 小件员所在公司编号
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetCpCode(_cpCode string) error {
    r._cpCode = _cpCode
    r.Set("cp_code", _cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetCpCode() string {
    return r._cpCode
}
// CpUserId Setter
// 小件员员工编号
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetCpUserId(_cpUserId string) error {
    r._cpUserId = _cpUserId
    r.Set("cp_user_id", _cpUserId)
    return nil
}

// CpUserId Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetCpUserId() string {
    return r._cpUserId
}
// TimeStamp Setter
// 上报时间，格式：yyyy-MM-dd HH:mm:ss
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetTimeStamp(_timeStamp string) error {
    r._timeStamp = _timeStamp
    r.Set("time_stamp", _timeStamp)
    return nil
}

// TimeStamp Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetTimeStamp() string {
    return r._timeStamp
}
// Source Setter
// 来源：1.小件员app sdk 2.驿站 3. 裹裹 10001.圆通行者
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetSource(_source int64) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetSource() int64 {
    return r._source
}
// Lng Setter
// 经度
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetLng(_lng string) error {
    r._lng = _lng
    r.Set("lng", _lng)
    return nil
}

// Lng Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetLng() string {
    return r._lng
}
// Lat Setter
// 纬度
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetLat(_lat string) error {
    r._lat = _lat
    r.Set("lat", _lat)
    return nil
}

// Lat Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetLat() string {
    return r._lat
}
// GpsType Setter
// 0 安卓定位，     1 苹果定位，  2 其他系统定位，   10 高德定位，  11 百度定位，  12 google定位     13 其他
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetGpsType(_gpsType string) error {
    r._gpsType = _gpsType
    r.Set("gps_type", _gpsType)
    return nil
}

// GpsType Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetGpsType() string {
    return r._gpsType
}
