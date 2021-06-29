package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件试妆镜数据回流 APIRequest
taobao.smartstore.device.makeupmirror.feedback

智慧门店试妆镜设备回流规则（适用于试妆镜等）</br>
1.回流的数据属于当前授权的用户，回流的设备device_code由当前应用添加</br>
2.对于快闪店的智能硬件不需要授权</br>
3.action为SKU_CLICK时，sku_id必须传入</br>
4.action为ITEM_CLICK、SAMPLE_CLICK、BUY_CLICK、ITEM_FAVOR时，item_id必须传入,且必须是淘宝商品的数字id </br>
5.skin_detection 和scalp_detection 涉及相关检测功能的硬件设备回传 </br>
6.每一个acion都必须传入用户操作时间op_time </br>
7.outer_biz_id 用于硬件设备大量数据回流场景，服务商本地日志统计系统对一条日志记录生成唯一标识。 平台后端会对传入的outer_biz_id 做去重处理</br>
*/
type TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest struct {
    model.Params

    // 肌肤检测结果，"1. Moisture 水份/     2. Sebum (U/T Zone) 油份(U/T区)/ 3. Pore 毛孔 4. Melanin 色斑 5. Acne (UV) 暗疮(紫外线)  6. Wrinkle 皱纹 7. Sensitivity 敏感度  数字指标：  行业平均值 industry_average  当前顾客数值 current_customer  检测结果 detection_result  检测描述 detection_description 数字指标可以为空，但是单个key请务必完整。如 ""detection_result"":""""  {  ""moisture"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""湿度高"",   ""detection_description"":""完美的肌肤特质""  },  ""sebum"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""混合型"",   ""detection_description"":""该顾客混合了过多的油脂""  },  ""pore"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""毛孔比较细小""  },  ""melanin"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""特别护理"",   ""detection_description"":""有较大的色斑区域""  },  ""wrinkle"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""特别护理"",   ""detection_description"":""有很多黑色粉刺的大毛孔""  },  ""acne"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""没有皱纹，皮肤良好""  },  ""sensitivity"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""有部分角质可以进行护理""  } }"
    skinDetection   string 

    // 头皮检查结果，"1. Alopecia  脱发状态     2. Scalp state 头皮状态 3. Hair density 头发密度 4. Scalp of scalp 头皮角质 5. Nudity of scalp blood vessels 头皮血管裸露  6. Hair thickness 头发厚度 7. Hair follicle state 毛囊状态  数字指标：  行业平均值 industry_average  当前顾客数值 current_customer  检测结果 detection_result  检测描述 detection_description 数字指标可以为空，但是单个key请务必完整。如 ""detection_result"":""""  {  ""alopecia"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""湿度高"",   ""detection_description"":""完美的肌肤特质""  },  ""state"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""混合型"",   ""detection_description"":""该顾客混合了过多的油脂""  },  ""hair_density"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""毛孔比较细小""  },  ""scalp"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""特别护理"",   ""detection_description"":""有较大的色斑区域""  },  ""nudity_blood"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""特别护理"",   ""detection_description"":""有很多黑色粉刺的大毛孔""  },  ""hair_thickness"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""没有皱纹，皮肤良好""  },  ""hair_follicle_state"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""有部分角质可以进行护理""  } }"
    scalpDetection   string 

    // 硬件CODE
    deviceCode   string 

    // 字段废弃
    endTime   string 

    // 字段废弃，考虑兼容，等同于op_time，两个必须传一个
    startTime   string 

    // ACTION枚举值： ITEM_CLICK（商品点击时，必须设置ITEM_ID） SKU_CLICK（SKU点击时，必须设置SKU_ID) THEME_MAKEUP_CLICK(点击推荐主题妆) SAMPLE_CLICK（点击领取小样，必须设置ITEM_ID） RECEIVE_COUPONS（领取优惠券时，必须设置COUPON_ID）  BUY_CLICK（点击购买，购买PV，必须设置ITEM_ID）  ITEM_FAVOR(商品收藏时，必须设置item_id) PHOTO_CLICK（拍摄照片） GET_PHOTO（获取照片，必须设置user_nick） SHARE_CLICK（点击分享）
    action   string 

    // 商品ID，item_id 在action为ITEM_CLICK、SAMPLE_CLICK、BUY_CLICK、ITEM_FAVOR必须传入， 必须使用淘宝商品id，否则失败
    itemId   string 

    // "例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
    couponId   string 

    // 数据外部编码，保证数据唯一性
    outerBizId   string 

    // 操作时间，后续统一使用该字段，考虑兼容，start_time跟该字段含义一致
    opTime   string 

    // 硬件识别的用户标识
    outerUser   string 

    // 库存ID，sku_id 在action为sku_click时必须传入； 必须使用淘宝商品id，否则失败。
    skuId   string 

    // 用户混淆nick
    userNick   string 

}

func NewTaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest() *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest{
    return &TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetApiMethodName() string {
    return "taobao.smartstore.device.makeupmirror.feedback"
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetSkinDetection(skinDetection string) error {
    r.skinDetection = skinDetection
    r.Set("skin_detection", skinDetection)
    return nil
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetSkinDetection() string {
    return r.skinDetection
}

func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetScalpDetection(scalpDetection string) error {
    r.scalpDetection = scalpDetection
    r.Set("scalp_detection", scalpDetection)
    return nil
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetScalpDetection() string {
    return r.scalpDetection
}

func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetDeviceCode() string {
    return r.deviceCode
}

func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetAction() string {
    return r.action
}

func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetItemId() string {
    return r.itemId
}

func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetCouponId(couponId string) error {
    r.couponId = couponId
    r.Set("coupon_id", couponId)
    return nil
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetCouponId() string {
    return r.couponId
}

func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetOuterBizId(outerBizId string) error {
    r.outerBizId = outerBizId
    r.Set("outer_biz_id", outerBizId)
    return nil
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetOuterBizId() string {
    return r.outerBizId
}

func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetOpTime(opTime string) error {
    r.opTime = opTime
    r.Set("op_time", opTime)
    return nil
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetOpTime() string {
    return r.opTime
}

func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetOuterUser(outerUser string) error {
    r.outerUser = outerUser
    r.Set("outer_user", outerUser)
    return nil
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetOuterUser() string {
    return r.outerUser
}

func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetSkuId(skuId string) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetSkuId() string {
    return r.skuId
}

func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetUserNick() string {
    return r.userNick
}

