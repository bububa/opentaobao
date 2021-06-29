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
    cpCode   string
    // 小件员员工编号
    cpUserId   string
    // 上报时间，格式：yyyy-MM-dd HH:mm:ss
    timeStamp   string
    // 来源：1.小件员app sdk 2.驿站 3. 裹裹 10001.圆通行者
    source   int64
    // 经度
    lng   string
    // 纬度
    lat   string
    // 0 安卓定位，     1 苹果定位，  2 其他系统定位，   10 高德定位，  11 百度定位，  12 google定位     13 其他
    gpsType   string
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
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetCpCode(cpCode string) error {
    r.cpCode = cpCode
    r.Set("cp_code", cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetCpCode() string {
    return r.cpCode
}
// CpUserId Setter
// 小件员员工编号
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetCpUserId(cpUserId string) error {
    r.cpUserId = cpUserId
    r.Set("cp_user_id", cpUserId)
    return nil
}

// CpUserId Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetCpUserId() string {
    return r.cpUserId
}
// TimeStamp Setter
// 上报时间，格式：yyyy-MM-dd HH:mm:ss
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetTimeStamp(timeStamp string) error {
    r.timeStamp = timeStamp
    r.Set("time_stamp", timeStamp)
    return nil
}

// TimeStamp Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetTimeStamp() string {
    return r.timeStamp
}
// Source Setter
// 来源：1.小件员app sdk 2.驿站 3. 裹裹 10001.圆通行者
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetSource(source int64) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetSource() int64 {
    return r.source
}
// Lng Setter
// 经度
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetLng(lng string) error {
    r.lng = lng
    r.Set("lng", lng)
    return nil
}

// Lng Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetLng() string {
    return r.lng
}
// Lat Setter
// 纬度
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetLat(lat string) error {
    r.lat = lat
    r.Set("lat", lat)
    return nil
}

// Lat Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetLat() string {
    return r.lat
}
// GpsType Setter
// 0 安卓定位，     1 苹果定位，  2 其他系统定位，   10 高德定位，  11 百度定位，  12 google定位     13 其他
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) SetGpsType(gpsType string) error {
    r.gpsType = gpsType
    r.Set("gps_type", gpsType)
    return nil
}

// GpsType Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest) GetGpsType() string {
    return r.gpsType
}
