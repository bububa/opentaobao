package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店货架商品列表信息查询 APIRequest
taobao.koubei.mall.common.item.shelf.page

查询口碑综合体内门店货架商品列表信息接口
*/
type TaobaoKoubeiMallCommonItemShelfPageRequest struct {
    model.Params

    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    dataSetId   string 

    // 商圈ID
    mallId   string 

    // 商圈内的门店ID
    storeId   string 

    // 分页查询起始值，默认为0
    start   int64 

    // 每页查询量，固定8个
    pageSize   int64 

    // 口碑城市编码（示例：杭州市330100）
    cityCode   string 

    // 经度（终端设备地理位置）
    longitude   string 

    // 纬度（终端设备地理位置）
    latitude   string 

    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    displayChannel   string 

    // 终端设备描述(中、英文均可)
    terminalType   string 

    // 支付宝/口碑/淘宝app版本号
    appVersion   string 

}

func NewTaobaoKoubeiMallCommonItemShelfPageRequest() *TaobaoKoubeiMallCommonItemShelfPageRequest{
    return &TaobaoKoubeiMallCommonItemShelfPageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.item.shelf.page"
}

func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetDataSetId() string {
    return r.dataSetId
}

func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetMallId() string {
    return r.mallId
}

func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetStoreId() string {
    return r.storeId
}

func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetStart() int64 {
    return r.start
}

func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetCityCode() string {
    return r.cityCode
}

func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetLongitude() string {
    return r.longitude
}

func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetLatitude() string {
    return r.latitude
}

func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetDisplayChannel(displayChannel string) error {
    r.displayChannel = displayChannel
    r.Set("display_channel", displayChannel)
    return nil
}

func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetDisplayChannel() string {
    return r.displayChannel
}

func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetTerminalType() string {
    return r.terminalType
}

func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetAppVersion() string {
    return r.appVersion
}

