package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商圈详细信息 APIRequest
taobao.koubei.mall.common.mall.detail.get

查询口碑综合体-商圈详细信息，包含商圈基础信息、门店类目分类、商圈推荐商品等模块信息
*/
type TaobaoKoubeiMallCommonMallDetailGetRequest struct {
    model.Params

    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    dataSetId   string 

    // 商圈ID
    mallId   string 

    // 经度（终端设备地理位置）
    longitude   string 

    // 纬度（终端设备地理位置）
    latitude   string 

    // 终端设备描述(中、英文均可)
    terminalType   string 

    // 支付宝/口碑/淘宝app版本号
    appVersion   string 

    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    displayChannel   string 

    // 口碑城市编码（示例：杭州市330100）
    cityCode   string 

}

func NewTaobaoKoubeiMallCommonMallDetailGetRequest() *TaobaoKoubeiMallCommonMallDetailGetRequest{
    return &TaobaoKoubeiMallCommonMallDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoKoubeiMallCommonMallDetailGetRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.mall.detail.get"
}

func (r TaobaoKoubeiMallCommonMallDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoKoubeiMallCommonMallDetailGetRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

func (r TaobaoKoubeiMallCommonMallDetailGetRequest) GetDataSetId() string {
    return r.dataSetId
}

func (r *TaobaoKoubeiMallCommonMallDetailGetRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

func (r TaobaoKoubeiMallCommonMallDetailGetRequest) GetMallId() string {
    return r.mallId
}

func (r *TaobaoKoubeiMallCommonMallDetailGetRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r TaobaoKoubeiMallCommonMallDetailGetRequest) GetLongitude() string {
    return r.longitude
}

func (r *TaobaoKoubeiMallCommonMallDetailGetRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r TaobaoKoubeiMallCommonMallDetailGetRequest) GetLatitude() string {
    return r.latitude
}

func (r *TaobaoKoubeiMallCommonMallDetailGetRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r TaobaoKoubeiMallCommonMallDetailGetRequest) GetTerminalType() string {
    return r.terminalType
}

func (r *TaobaoKoubeiMallCommonMallDetailGetRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

func (r TaobaoKoubeiMallCommonMallDetailGetRequest) GetAppVersion() string {
    return r.appVersion
}

func (r *TaobaoKoubeiMallCommonMallDetailGetRequest) SetDisplayChannel(displayChannel string) error {
    r.displayChannel = displayChannel
    r.Set("display_channel", displayChannel)
    return nil
}

func (r TaobaoKoubeiMallCommonMallDetailGetRequest) GetDisplayChannel() string {
    return r.displayChannel
}

func (r *TaobaoKoubeiMallCommonMallDetailGetRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r TaobaoKoubeiMallCommonMallDetailGetRequest) GetCityCode() string {
    return r.cityCode
}

