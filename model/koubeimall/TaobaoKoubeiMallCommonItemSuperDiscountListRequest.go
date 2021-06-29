package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商圈内的超值特惠商品信息 APIRequest
taobao.koubei.mall.common.item.super.discount.list

查询商圈超值特惠商品信息列表
*/
type TaobaoKoubeiMallCommonItemSuperDiscountListRequest struct {
    model.Params

    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    dataSetId   string 

    // 商圈ID
    mallId   string 

    // 查询商品最大个数，最大值50
    itemSize   int64 

    // 经度（终端设备地理位置）
    longitude   string 

    // 纬度（终端设备地理位置）
    latitude   string 

    // 口碑城市编码（示例：杭州市330100）
    cityCode   string 

    // 终端设备描述(中、英文均可)
    terminalType   string 

    // 支付宝/口碑/淘宝app版本号
    appVersion   string 

    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    displayChannel   string 

}

func NewTaobaoKoubeiMallCommonItemSuperDiscountListRequest() *TaobaoKoubeiMallCommonItemSuperDiscountListRequest{
    return &TaobaoKoubeiMallCommonItemSuperDiscountListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.item.super.discount.list"
}

func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetDataSetId() string {
    return r.dataSetId
}

func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetMallId() string {
    return r.mallId
}

func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetItemSize(itemSize int64) error {
    r.itemSize = itemSize
    r.Set("item_size", itemSize)
    return nil
}

func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetItemSize() int64 {
    return r.itemSize
}

func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetLongitude() string {
    return r.longitude
}

func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetLatitude() string {
    return r.latitude
}

func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetCityCode() string {
    return r.cityCode
}

func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetTerminalType() string {
    return r.terminalType
}

func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetAppVersion() string {
    return r.appVersion
}

func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetDisplayChannel(displayChannel string) error {
    r.displayChannel = displayChannel
    r.Set("display_channel", displayChannel)
    return nil
}

func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetDisplayChannel() string {
    return r.displayChannel
}

