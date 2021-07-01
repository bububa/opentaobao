package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件互动大屏数据回流 API请求
taobao.smartstore.device.storegame.feedback

智慧门店互动引流屏设备回流规则：（适用于智慧迎宾屏、互动游戏、互动拍照、娃娃机等）</br>
1.回流的数据属于当前授权的用户，回流的设备device_code由当前应用添加</br>
2.对于快闪店的智能硬件不需要授权</br>
3.每一个action都必须传入用户操作时间op_time；（start/end_time后续废弃）</br>
4.action为WINNING_PRIZE时，需传入draw_result，只能传入0或者1</br>
5.outer_biz_id 用于硬件设备大量数据回流场景，服务商本地日志统计系统对一条日志记录生成唯一标识。 平台后端会对传入的outer_biz_id 做去重处理</br>
6.outer_user 用于标识不能获取淘宝账号的游客</br>
*/
type TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest struct {
    model.Params
    // 商品ID，item_id 在action为ITEM_CLICK时必须传入。 必须使用淘宝商品id，否则失败。
    _itemId   string
    // 游戏名称
    _gameName   string
    // 硬件CODE
    _deviceCode   string
    // 字段废弃
    _endTime   string
    // 字段废弃，考虑兼容，等同于op_time，两个必须传一个
    _startTime   string
    // ACTION枚举值：  BODY_SENSOR（通过人体感应、人脸识别成功识别到人） PHOTO_CLICK（用户在屏幕拍摄照片） GET_PHOTO（用户扫码获取照片，必须设置user_nick） ITEM_CLICK（商品点击时必须设置ITEM_ID）  GAME_START（开始游戏，可以 设置 user）  GAME_OVER_WITHOUT_PROMOTION（游戏结束）  WINNING_PRIZE（中奖，必须设置DRAW_RESULT） SHARE_CLICK（点击分享） RECEIVE_COUPONS (扫码领取优惠券时必须设置COUPON_ID)
    _action   string
    // 有则传入，没有可以不传。"例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
    _couponId   string
    // 用户昵称，混淆的淘宝NIck
    _userNick   string
    // 抽奖结果 ，如果传入，0：表示没中奖，1：表示中奖。该值必须是0或者1，传入其他失败。
    _drawResult   string
    // 数据外部编码，保证数据唯一性
    _outerBizId   string
    // 操作时间，后续统一使用该字段，考虑兼容，start_time跟该字段含义一致
    _opTime   string
    // 硬件识别的用户标识
    _outerUser   string
}

// 初始化TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest对象
func NewTaobaoSmartstoreDeviceStoregameFeedbackRequest() *TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest{
    return &TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetApiMethodName() string {
    return "taobao.smartstore.device.storegame.feedback"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID，item_id 在action为ITEM_CLICK时必须传入。 必须使用淘宝商品id，否则失败。
func (r *TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) SetItemId(_itemId string) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetItemId() string {
    return r._itemId
}
// GameName Setter
// 游戏名称
func (r *TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) SetGameName(_gameName string) error {
    r._gameName = _gameName
    r.Set("game_name", _gameName)
    return nil
}

// GameName Getter
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetGameName() string {
    return r._gameName
}
// DeviceCode Setter
// 硬件CODE
func (r *TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetDeviceCode() string {
    return r._deviceCode
}
// EndTime Setter
// 字段废弃
func (r *TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetEndTime() string {
    return r._endTime
}
// StartTime Setter
// 字段废弃，考虑兼容，等同于op_time，两个必须传一个
func (r *TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetStartTime() string {
    return r._startTime
}
// Action Setter
// ACTION枚举值：  BODY_SENSOR（通过人体感应、人脸识别成功识别到人） PHOTO_CLICK（用户在屏幕拍摄照片） GET_PHOTO（用户扫码获取照片，必须设置user_nick） ITEM_CLICK（商品点击时必须设置ITEM_ID）  GAME_START（开始游戏，可以 设置 user）  GAME_OVER_WITHOUT_PROMOTION（游戏结束）  WINNING_PRIZE（中奖，必须设置DRAW_RESULT） SHARE_CLICK（点击分享） RECEIVE_COUPONS (扫码领取优惠券时必须设置COUPON_ID)
func (r *TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetAction() string {
    return r._action
}
// CouponId Setter
// 有则传入，没有可以不传。"例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
func (r *TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) SetCouponId(_couponId string) error {
    r._couponId = _couponId
    r.Set("coupon_id", _couponId)
    return nil
}

// CouponId Getter
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetCouponId() string {
    return r._couponId
}
// UserNick Setter
// 用户昵称，混淆的淘宝NIck
func (r *TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetUserNick() string {
    return r._userNick
}
// DrawResult Setter
// 抽奖结果 ，如果传入，0：表示没中奖，1：表示中奖。该值必须是0或者1，传入其他失败。
func (r *TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) SetDrawResult(_drawResult string) error {
    r._drawResult = _drawResult
    r.Set("draw_result", _drawResult)
    return nil
}

// DrawResult Getter
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetDrawResult() string {
    return r._drawResult
}
// OuterBizId Setter
// 数据外部编码，保证数据唯一性
func (r *TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) SetOuterBizId(_outerBizId string) error {
    r._outerBizId = _outerBizId
    r.Set("outer_biz_id", _outerBizId)
    return nil
}

// OuterBizId Getter
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetOuterBizId() string {
    return r._outerBizId
}
// OpTime Setter
// 操作时间，后续统一使用该字段，考虑兼容，start_time跟该字段含义一致
func (r *TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) SetOpTime(_opTime string) error {
    r._opTime = _opTime
    r.Set("op_time", _opTime)
    return nil
}

// OpTime Getter
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetOpTime() string {
    return r._opTime
}
// OuterUser Setter
// 硬件识别的用户标识
func (r *TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) SetOuterUser(_outerUser string) error {
    r._outerUser = _outerUser
    r.Set("outer_user", _outerUser)
    return nil
}

// OuterUser Getter
func (r TaobaoSmartstoreDeviceStoregameFeedbackAPIRequest) GetOuterUser() string {
    return r._outerUser
}
