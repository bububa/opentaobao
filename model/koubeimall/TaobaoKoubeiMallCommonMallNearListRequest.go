package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据POI查询附近商圈列表信息 APIRequest
taobao.koubei.mall.common.mall.near.list

通过用户/终端设备地理位置POI信息，查询附近商圈信息
*/
type TaobaoKoubeiMallCommonMallNearListRequest struct {
    model.Params

    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    dataSetId   string 

    // 召回半径，单位m，最大数值不能超过10km（该字段为空，默认全城召回）
    radius   int64 

    // 查询个数，最大查询量不能超过50个
    size   int64 

    // 经度（终端设备地理位置）
    longitude   string 

    // 纬度（终端设备地理位置）
    latitude   string 

    // 口碑城市编码（示例：杭州市330100）
    cityCode   string 

    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    displayChannel   string 

    // 支付宝/口碑/淘宝app版本号
    appVersion   string 

    // 终端设备描述(中、英文均可)
    terminalType   string 

}

func NewTaobaoKoubeiMallCommonMallNearListRequest() *TaobaoKoubeiMallCommonMallNearListRequest{
    return &TaobaoKoubeiMallCommonMallNearListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoKoubeiMallCommonMallNearListRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.mall.near.list"
}

func (r TaobaoKoubeiMallCommonMallNearListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

func (r TaobaoKoubeiMallCommonMallNearListRequest) GetDataSetId() string {
    return r.dataSetId
}

func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetRadius(radius int64) error {
    r.radius = radius
    r.Set("radius", radius)
    return nil
}

func (r TaobaoKoubeiMallCommonMallNearListRequest) GetRadius() int64 {
    return r.radius
}

func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetSize(size int64) error {
    r.size = size
    r.Set("size", size)
    return nil
}

func (r TaobaoKoubeiMallCommonMallNearListRequest) GetSize() int64 {
    return r.size
}

func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r TaobaoKoubeiMallCommonMallNearListRequest) GetLongitude() string {
    return r.longitude
}

func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r TaobaoKoubeiMallCommonMallNearListRequest) GetLatitude() string {
    return r.latitude
}

func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r TaobaoKoubeiMallCommonMallNearListRequest) GetCityCode() string {
    return r.cityCode
}

func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetDisplayChannel(displayChannel string) error {
    r.displayChannel = displayChannel
    r.Set("display_channel", displayChannel)
    return nil
}

func (r TaobaoKoubeiMallCommonMallNearListRequest) GetDisplayChannel() string {
    return r.displayChannel
}

func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

func (r TaobaoKoubeiMallCommonMallNearListRequest) GetAppVersion() string {
    return r.appVersion
}

func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r TaobaoKoubeiMallCommonMallNearListRequest) GetTerminalType() string {
    return r.terminalType
}

