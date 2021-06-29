package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件试妆镜数据回流 API请求
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
    _skinDetection   string
    // 头皮检查结果，"1. Alopecia  脱发状态     2. Scalp state 头皮状态 3. Hair density 头发密度 4. Scalp of scalp 头皮角质 5. Nudity of scalp blood vessels 头皮血管裸露  6. Hair thickness 头发厚度 7. Hair follicle state 毛囊状态  数字指标：  行业平均值 industry_average  当前顾客数值 current_customer  检测结果 detection_result  检测描述 detection_description 数字指标可以为空，但是单个key请务必完整。如 ""detection_result"":""""  {  ""alopecia"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""湿度高"",   ""detection_description"":""完美的肌肤特质""  },  ""state"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""混合型"",   ""detection_description"":""该顾客混合了过多的油脂""  },  ""hair_density"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""毛孔比较细小""  },  ""scalp"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""特别护理"",   ""detection_description"":""有较大的色斑区域""  },  ""nudity_blood"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""特别护理"",   ""detection_description"":""有很多黑色粉刺的大毛孔""  },  ""hair_thickness"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""没有皱纹，皮肤良好""  },  ""hair_follicle_state"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""有部分角质可以进行护理""  } }"
    _scalpDetection   string
    // 硬件CODE
    _deviceCode   string
    // 字段废弃
    _endTime   string
    // 字段废弃，考虑兼容，等同于op_time，两个必须传一个
    _startTime   string
    // ACTION枚举值： ITEM_CLICK（商品点击时，必须设置ITEM_ID） SKU_CLICK（SKU点击时，必须设置SKU_ID) THEME_MAKEUP_CLICK(点击推荐主题妆) SAMPLE_CLICK（点击领取小样，必须设置ITEM_ID） RECEIVE_COUPONS（领取优惠券时，必须设置COUPON_ID）  BUY_CLICK（点击购买，购买PV，必须设置ITEM_ID）  ITEM_FAVOR(商品收藏时，必须设置item_id) PHOTO_CLICK（拍摄照片） GET_PHOTO（获取照片，必须设置user_nick） SHARE_CLICK（点击分享）
    _action   string
    // 商品ID，item_id 在action为ITEM_CLICK、SAMPLE_CLICK、BUY_CLICK、ITEM_FAVOR必须传入， 必须使用淘宝商品id，否则失败
    _itemId   string
    // "例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
    _couponId   string
    // 数据外部编码，保证数据唯一性
    _outerBizId   string
    // 操作时间，后续统一使用该字段，考虑兼容，start_time跟该字段含义一致
    _opTime   string
    // 硬件识别的用户标识
    _outerUser   string
    // 库存ID，sku_id 在action为sku_click时必须传入； 必须使用淘宝商品id，否则失败。
    _skuId   string
    // 用户混淆nick
    _userNick   string
}

// 初始化TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest对象
func NewTaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest() *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest{
    return &TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetApiMethodName() string {
    return "taobao.smartstore.device.makeupmirror.feedback"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkinDetection Setter
// 肌肤检测结果，"1. Moisture 水份/     2. Sebum (U/T Zone) 油份(U/T区)/ 3. Pore 毛孔 4. Melanin 色斑 5. Acne (UV) 暗疮(紫外线)  6. Wrinkle 皱纹 7. Sensitivity 敏感度  数字指标：  行业平均值 industry_average  当前顾客数值 current_customer  检测结果 detection_result  检测描述 detection_description 数字指标可以为空，但是单个key请务必完整。如 ""detection_result"":""""  {  ""moisture"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""湿度高"",   ""detection_description"":""完美的肌肤特质""  },  ""sebum"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""混合型"",   ""detection_description"":""该顾客混合了过多的油脂""  },  ""pore"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""毛孔比较细小""  },  ""melanin"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""特别护理"",   ""detection_description"":""有较大的色斑区域""  },  ""wrinkle"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""特别护理"",   ""detection_description"":""有很多黑色粉刺的大毛孔""  },  ""acne"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""没有皱纹，皮肤良好""  },  ""sensitivity"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""有部分角质可以进行护理""  } }"
func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetSkinDetection(_skinDetection string) error {
    r._skinDetection = _skinDetection
    r.Set("skin_detection", _skinDetection)
    return nil
}

// SkinDetection Getter
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetSkinDetection() string {
    return r._skinDetection
}
// ScalpDetection Setter
// 头皮检查结果，"1. Alopecia  脱发状态     2. Scalp state 头皮状态 3. Hair density 头发密度 4. Scalp of scalp 头皮角质 5. Nudity of scalp blood vessels 头皮血管裸露  6. Hair thickness 头发厚度 7. Hair follicle state 毛囊状态  数字指标：  行业平均值 industry_average  当前顾客数值 current_customer  检测结果 detection_result  检测描述 detection_description 数字指标可以为空，但是单个key请务必完整。如 ""detection_result"":""""  {  ""alopecia"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""湿度高"",   ""detection_description"":""完美的肌肤特质""  },  ""state"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""混合型"",   ""detection_description"":""该顾客混合了过多的油脂""  },  ""hair_density"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""毛孔比较细小""  },  ""scalp"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""特别护理"",   ""detection_description"":""有较大的色斑区域""  },  ""nudity_blood"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""特别护理"",   ""detection_description"":""有很多黑色粉刺的大毛孔""  },  ""hair_thickness"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""没有皱纹，皮肤良好""  },  ""hair_follicle_state"": {   ""industry_average"": ""26"",   ""current_customer"": ""39"",   ""detection_result"":""良好"",   ""detection_description"":""有部分角质可以进行护理""  } }"
func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetScalpDetection(_scalpDetection string) error {
    r._scalpDetection = _scalpDetection
    r.Set("scalp_detection", _scalpDetection)
    return nil
}

// ScalpDetection Getter
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetScalpDetection() string {
    return r._scalpDetection
}
// DeviceCode Setter
// 硬件CODE
func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetDeviceCode() string {
    return r._deviceCode
}
// EndTime Setter
// 字段废弃
func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetEndTime() string {
    return r._endTime
}
// StartTime Setter
// 字段废弃，考虑兼容，等同于op_time，两个必须传一个
func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetStartTime() string {
    return r._startTime
}
// Action Setter
// ACTION枚举值： ITEM_CLICK（商品点击时，必须设置ITEM_ID） SKU_CLICK（SKU点击时，必须设置SKU_ID) THEME_MAKEUP_CLICK(点击推荐主题妆) SAMPLE_CLICK（点击领取小样，必须设置ITEM_ID） RECEIVE_COUPONS（领取优惠券时，必须设置COUPON_ID）  BUY_CLICK（点击购买，购买PV，必须设置ITEM_ID）  ITEM_FAVOR(商品收藏时，必须设置item_id) PHOTO_CLICK（拍摄照片） GET_PHOTO（获取照片，必须设置user_nick） SHARE_CLICK（点击分享）
func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetAction() string {
    return r._action
}
// ItemId Setter
// 商品ID，item_id 在action为ITEM_CLICK、SAMPLE_CLICK、BUY_CLICK、ITEM_FAVOR必须传入， 必须使用淘宝商品id，否则失败
func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetItemId(_itemId string) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetItemId() string {
    return r._itemId
}
// CouponId Setter
// "例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetCouponId(_couponId string) error {
    r._couponId = _couponId
    r.Set("coupon_id", _couponId)
    return nil
}

// CouponId Getter
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetCouponId() string {
    return r._couponId
}
// OuterBizId Setter
// 数据外部编码，保证数据唯一性
func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetOuterBizId(_outerBizId string) error {
    r._outerBizId = _outerBizId
    r.Set("outer_biz_id", _outerBizId)
    return nil
}

// OuterBizId Getter
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetOuterBizId() string {
    return r._outerBizId
}
// OpTime Setter
// 操作时间，后续统一使用该字段，考虑兼容，start_time跟该字段含义一致
func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetOpTime(_opTime string) error {
    r._opTime = _opTime
    r.Set("op_time", _opTime)
    return nil
}

// OpTime Getter
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetOpTime() string {
    return r._opTime
}
// OuterUser Setter
// 硬件识别的用户标识
func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetOuterUser(_outerUser string) error {
    r._outerUser = _outerUser
    r.Set("outer_user", _outerUser)
    return nil
}

// OuterUser Getter
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetOuterUser() string {
    return r._outerUser
}
// SkuId Setter
// 库存ID，sku_id 在action为sku_click时必须传入； 必须使用淘宝商品id，否则失败。
func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetSkuId(_skuId string) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetSkuId() string {
    return r._skuId
}
// UserNick Setter
// 用户混淆nick
func (r *TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r TaobaoSmartstoreDeviceMakeupmirrorFeedbackRequest) GetUserNick() string {
    return r._userNick
}
