package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件云货架数据回流 API请求
taobao.smartstore.device.iashelf.feedback

智慧门店云货架设备回流规则：（互动云货架、VR云货架通用）</br>
1.回流的数据属于当前授权的用户，回流的设备device_code由当前应用添加</br>
2.对于快闪店的智能硬件不需要授权</br>
3.每一个action都必须传入用户操作时间op_time（start/end_time后续废弃）</br>
4.action为ITEM_CLICK、ITEM_SENSOR、BUY_CLICK时，item_id必须传入,且必须是淘宝商品的数字id</br>
5.outer_biz_id 用于硬件设备大量数据回流场景，服务商本地日志统计系统对一条日志记录生成唯一标识。 平台后端会对传入的outer_biz_id 做去重处理</br>
6.outer_user 用于标识不能获取淘宝账号的游客</br>
*/
type TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest struct {
    model.Params
    // 硬件CODE
    _deviceCode   string
    // 字段废弃
    _endTime   string
    // 字段废弃，考虑兼容，等同于op_time，两个必须传一个
    _startTime   string
    // ACTION枚举值： BODY_SENSOR（通过人体感应、人脸识别成功识别到人） ITEM_SENSOR（通过RFID、蓝牙等感应设备识别商品，必须设置ITEM_ID） ITEM_CLICK（商品点击时必须设置ITEM_ID）  RECEIVE_COUPONS（扫码领取优惠券时必须设置COUPON_ID） BUY_CLICK（点击购买，必须设置ITEM_ID）  SHARE_CLICK（点击分享）
    _action   string
    // 商品ID，item_id 在action为ITEM_CLICK、ITEM_SENSOR、BUY_CLICK必须传入；  必须使用淘宝商品id，否则失败。
    _itemId   string
    // "例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
    _couponId   string
    // 数据外部编码，保证数据唯一性
    _outerBizId   string
    // 操作时间，后续统一使用该字段，考虑兼容，start_time跟该字段含义一致
    _opTime   string
    // 用户昵称（混淆）
    _userNick   string
    // 硬件识别的用户标识
    _outerUser   string
}

// 初始化TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest对象
func NewTaobaoSmartstoreDeviceIashelfFeedbackRequest() *TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest{
    return &TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) GetApiMethodName() string {
    return "taobao.smartstore.device.iashelf.feedback"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 硬件CODE
func (r *TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) GetDeviceCode() string {
    return r._deviceCode
}
// EndTime Setter
// 字段废弃
func (r *TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) GetEndTime() string {
    return r._endTime
}
// StartTime Setter
// 字段废弃，考虑兼容，等同于op_time，两个必须传一个
func (r *TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) GetStartTime() string {
    return r._startTime
}
// Action Setter
// ACTION枚举值： BODY_SENSOR（通过人体感应、人脸识别成功识别到人） ITEM_SENSOR（通过RFID、蓝牙等感应设备识别商品，必须设置ITEM_ID） ITEM_CLICK（商品点击时必须设置ITEM_ID）  RECEIVE_COUPONS（扫码领取优惠券时必须设置COUPON_ID） BUY_CLICK（点击购买，必须设置ITEM_ID）  SHARE_CLICK（点击分享）
func (r *TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) GetAction() string {
    return r._action
}
// ItemId Setter
// 商品ID，item_id 在action为ITEM_CLICK、ITEM_SENSOR、BUY_CLICK必须传入；  必须使用淘宝商品id，否则失败。
func (r *TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) SetItemId(_itemId string) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) GetItemId() string {
    return r._itemId
}
// CouponId Setter
// "例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
func (r *TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) SetCouponId(_couponId string) error {
    r._couponId = _couponId
    r.Set("coupon_id", _couponId)
    return nil
}

// CouponId Getter
func (r TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) GetCouponId() string {
    return r._couponId
}
// OuterBizId Setter
// 数据外部编码，保证数据唯一性
func (r *TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) SetOuterBizId(_outerBizId string) error {
    r._outerBizId = _outerBizId
    r.Set("outer_biz_id", _outerBizId)
    return nil
}

// OuterBizId Getter
func (r TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) GetOuterBizId() string {
    return r._outerBizId
}
// OpTime Setter
// 操作时间，后续统一使用该字段，考虑兼容，start_time跟该字段含义一致
func (r *TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) SetOpTime(_opTime string) error {
    r._opTime = _opTime
    r.Set("op_time", _opTime)
    return nil
}

// OpTime Getter
func (r TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) GetOpTime() string {
    return r._opTime
}
// UserNick Setter
// 用户昵称（混淆）
func (r *TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) GetUserNick() string {
    return r._userNick
}
// OuterUser Setter
// 硬件识别的用户标识
func (r *TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) SetOuterUser(_outerUser string) error {
    r._outerUser = _outerUser
    r.Set("outer_user", _outerUser)
    return nil
}

// OuterUser Getter
func (r TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest) GetOuterUser() string {
    return r._outerUser
}
