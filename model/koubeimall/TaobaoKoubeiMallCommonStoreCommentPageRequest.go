package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询门店评论详情信息 APIRequest
taobao.koubei.mall.common.store.comment.page

查询口碑综合体内的门店评论信息
*/
type TaobaoKoubeiMallCommonStoreCommentPageRequest struct {
    model.Params

    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    dataSetId   string 

    // 商圈ID
    mallId   string 

    // 门店ID
    storeId   string 

    // 查询起始值，默认为0
    start   int64 

    // 每页查询量，默认为20，最大数值20
    pageSize   int64 

    // 口碑城市编码（示例：杭州市330100）
    cityCode   string 

    // 纬度（终端设备地理位置）
    longitude   string 

    // 经度（终端设备地理位置）
    latitude   string 

    // 终端设备描述(中、英文均可)
    terminalType   string 

    // 支付宝/口碑/淘宝app版本号
    appVersion   string 

    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    displayChannel   string 

}

func NewTaobaoKoubeiMallCommonStoreCommentPageRequest() *TaobaoKoubeiMallCommonStoreCommentPageRequest{
    return &TaobaoKoubeiMallCommonStoreCommentPageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.store.comment.page"
}

func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetDataSetId() string {
    return r.dataSetId
}

func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetMallId() string {
    return r.mallId
}

func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetStoreId() string {
    return r.storeId
}

func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetStart() int64 {
    return r.start
}

func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetCityCode() string {
    return r.cityCode
}

func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetLongitude() string {
    return r.longitude
}

func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetLatitude() string {
    return r.latitude
}

func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetTerminalType() string {
    return r.terminalType
}

func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetAppVersion() string {
    return r.appVersion
}

func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetDisplayChannel(displayChannel string) error {
    r.displayChannel = displayChannel
    r.Set("display_channel", displayChannel)
    return nil
}

func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetDisplayChannel() string {
    return r.displayChannel
}

