package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品详情信息 APIRequest
taobao.koubei.mall.common.item.detail.query

查询口碑综合体内商品详情信息
*/
type TaobaoKoubeiMallCommonItemDetailQueryRequest struct {
    model.Params

    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    dataSetId   string 

    // 商圈ID
    mallId   string 

    // 门店ID
    storeId   string 

    // 商品ID
    itemId   string 

    // 口碑城市编码（示例：杭州市330100）
    cityCode   string 

    // 经度（终端设备地理位置）
    longitude   string 

    // 纬度（终端设备地理位置）
    latitude   string 

    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    displayChannel   string 

    // 支付宝/口碑/淘宝app版本号
    appVersion   string 

    // 终端设备描述(中、英文均可)
    terminalType   string 

}

func NewTaobaoKoubeiMallCommonItemDetailQueryRequest() *TaobaoKoubeiMallCommonItemDetailQueryRequest{
    return &TaobaoKoubeiMallCommonItemDetailQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoKoubeiMallCommonItemDetailQueryRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.item.detail.query"
}

func (r TaobaoKoubeiMallCommonItemDetailQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoKoubeiMallCommonItemDetailQueryRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

func (r TaobaoKoubeiMallCommonItemDetailQueryRequest) GetDataSetId() string {
    return r.dataSetId
}

func (r *TaobaoKoubeiMallCommonItemDetailQueryRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

func (r TaobaoKoubeiMallCommonItemDetailQueryRequest) GetMallId() string {
    return r.mallId
}

func (r *TaobaoKoubeiMallCommonItemDetailQueryRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoKoubeiMallCommonItemDetailQueryRequest) GetStoreId() string {
    return r.storeId
}

func (r *TaobaoKoubeiMallCommonItemDetailQueryRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoKoubeiMallCommonItemDetailQueryRequest) GetItemId() string {
    return r.itemId
}

func (r *TaobaoKoubeiMallCommonItemDetailQueryRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r TaobaoKoubeiMallCommonItemDetailQueryRequest) GetCityCode() string {
    return r.cityCode
}

func (r *TaobaoKoubeiMallCommonItemDetailQueryRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r TaobaoKoubeiMallCommonItemDetailQueryRequest) GetLongitude() string {
    return r.longitude
}

func (r *TaobaoKoubeiMallCommonItemDetailQueryRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r TaobaoKoubeiMallCommonItemDetailQueryRequest) GetLatitude() string {
    return r.latitude
}

func (r *TaobaoKoubeiMallCommonItemDetailQueryRequest) SetDisplayChannel(displayChannel string) error {
    r.displayChannel = displayChannel
    r.Set("display_channel", displayChannel)
    return nil
}

func (r TaobaoKoubeiMallCommonItemDetailQueryRequest) GetDisplayChannel() string {
    return r.displayChannel
}

func (r *TaobaoKoubeiMallCommonItemDetailQueryRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

func (r TaobaoKoubeiMallCommonItemDetailQueryRequest) GetAppVersion() string {
    return r.appVersion
}

func (r *TaobaoKoubeiMallCommonItemDetailQueryRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r TaobaoKoubeiMallCommonItemDetailQueryRequest) GetTerminalType() string {
    return r.terminalType
}

