package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-物料精选 APIRequest
taobao.tbk.dg.optimus.material

支持入参对应的“推广位”和官方提供的“物料id”，获取指定物料信息和推广链接，还可入参用户信息提供智能推荐（需智能推荐请先前协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
*/
type TaobaoTbkDgOptimusMaterialRequest struct {
    model.Params

    // 页大小，默认20，1~100
    pageSize   int64 

    // mm_xxx_xxx_xxx的第三位
    adzoneId   int64 

    // 第几页，默认：1
    pageNo   int64 

    // 官方的物料Id(详细物料id见：https://market.m.taobao.com/app/qn/toutiao-new/index-pc.html#/detail/10628875?_k=gpov9a)
    materialId   int64 

    // 智能匹配-设备号加密后的值（MD5加密需32位小写），类型为OAID时传原始OAID值
    deviceValue   string 

    // 智能匹配-设备号加密类型：MD5，类型为OAID时不传
    deviceEncrypt   string 

    // 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
    deviceType   string 

    // 内容专用-内容详情ID
    contentId   int64 

    // 内容专用-内容渠道信息
    contentSource   string 

    // 商品ID，用于相似商品推荐
    itemId   int64 

    // 选品库投放id
    favoritesId   string 

}

func NewTaobaoTbkDgOptimusMaterialRequest() *TaobaoTbkDgOptimusMaterialRequest{
    return &TaobaoTbkDgOptimusMaterialRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkDgOptimusMaterialRequest) GetApiMethodName() string {
    return "taobao.tbk.dg.optimus.material"
}

func (r TaobaoTbkDgOptimusMaterialRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkDgOptimusMaterialRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoTbkDgOptimusMaterialRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoTbkDgOptimusMaterialRequest) SetAdzoneId(adzoneId int64) error {
    r.adzoneId = adzoneId
    r.Set("adzone_id", adzoneId)
    return nil
}

func (r TaobaoTbkDgOptimusMaterialRequest) GetAdzoneId() int64 {
    return r.adzoneId
}

func (r *TaobaoTbkDgOptimusMaterialRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoTbkDgOptimusMaterialRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoTbkDgOptimusMaterialRequest) SetMaterialId(materialId int64) error {
    r.materialId = materialId
    r.Set("material_id", materialId)
    return nil
}

func (r TaobaoTbkDgOptimusMaterialRequest) GetMaterialId() int64 {
    return r.materialId
}

func (r *TaobaoTbkDgOptimusMaterialRequest) SetDeviceValue(deviceValue string) error {
    r.deviceValue = deviceValue
    r.Set("device_value", deviceValue)
    return nil
}

func (r TaobaoTbkDgOptimusMaterialRequest) GetDeviceValue() string {
    return r.deviceValue
}

func (r *TaobaoTbkDgOptimusMaterialRequest) SetDeviceEncrypt(deviceEncrypt string) error {
    r.deviceEncrypt = deviceEncrypt
    r.Set("device_encrypt", deviceEncrypt)
    return nil
}

func (r TaobaoTbkDgOptimusMaterialRequest) GetDeviceEncrypt() string {
    return r.deviceEncrypt
}

func (r *TaobaoTbkDgOptimusMaterialRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

func (r TaobaoTbkDgOptimusMaterialRequest) GetDeviceType() string {
    return r.deviceType
}

func (r *TaobaoTbkDgOptimusMaterialRequest) SetContentId(contentId int64) error {
    r.contentId = contentId
    r.Set("content_id", contentId)
    return nil
}

func (r TaobaoTbkDgOptimusMaterialRequest) GetContentId() int64 {
    return r.contentId
}

func (r *TaobaoTbkDgOptimusMaterialRequest) SetContentSource(contentSource string) error {
    r.contentSource = contentSource
    r.Set("content_source", contentSource)
    return nil
}

func (r TaobaoTbkDgOptimusMaterialRequest) GetContentSource() string {
    return r.contentSource
}

func (r *TaobaoTbkDgOptimusMaterialRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoTbkDgOptimusMaterialRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoTbkDgOptimusMaterialRequest) SetFavoritesId(favoritesId string) error {
    r.favoritesId = favoritesId
    r.Set("favorites_id", favoritesId)
    return nil
}

func (r TaobaoTbkDgOptimusMaterialRequest) GetFavoritesId() string {
    return r.favoritesId
}

