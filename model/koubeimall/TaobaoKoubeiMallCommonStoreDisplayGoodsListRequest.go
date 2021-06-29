package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询门店推荐菜信息 APIRequest
taobao.koubei.mall.common.store.display.goods.list

提供查询口碑商圈内的门店推荐菜信息
*/
type TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest struct {
    model.Params

    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    dataSetId   string 

    // 门店ID
    storeId   string 

    // 商圈ID
    mallId   string 

    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    displayChannel   string 

    // 口碑城市编码（示例：杭州市330100）
    cityCode   string 

    // 经度（终端设备地理位置）
    longitude   string 

    // 纬度（终端设备地理位置）
    latitude   string 

    // 终端设备描述(中、英文均可)
    terminalType   string 

    // 支付宝/口碑/淘宝app版本号
    appVersion   string 

}

func NewTaobaoKoubeiMallCommonStoreDisplayGoodsListRequest() *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest{
    return &TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.store.display.goods.list"
}

func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetDataSetId() string {
    return r.dataSetId
}

func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetStoreId() string {
    return r.storeId
}

func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetMallId() string {
    return r.mallId
}

func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetDisplayChannel(displayChannel string) error {
    r.displayChannel = displayChannel
    r.Set("display_channel", displayChannel)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetDisplayChannel() string {
    return r.displayChannel
}

func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetCityCode() string {
    return r.cityCode
}

func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetLongitude() string {
    return r.longitude
}

func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetLatitude() string {
    return r.latitude
}

func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetTerminalType() string {
    return r.terminalType
}

func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetAppVersion() string {
    return r.appVersion
}

