package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件互动大屏数据回流 APIRequest
taobao.smartstore.device.storegame.feedback

智慧门店互动引流屏设备回流规则：（适用于智慧迎宾屏、互动游戏、互动拍照、娃娃机等）</br>
1.回流的数据属于当前授权的用户，回流的设备device_code由当前应用添加</br>
2.对于快闪店的智能硬件不需要授权</br>
3.每一个action都必须传入用户操作时间op_time；（start/end_time后续废弃）</br>
4.action为WINNING_PRIZE时，需传入draw_result，只能传入0或者1</br>
5.outer_biz_id 用于硬件设备大量数据回流场景，服务商本地日志统计系统对一条日志记录生成唯一标识。 平台后端会对传入的outer_biz_id 做去重处理</br>
6.outer_user 用于标识不能获取淘宝账号的游客</br>
*/
type TaobaoSmartstoreDeviceStoregameFeedbackRequest struct {
    model.Params

    // 商品ID，item_id 在action为ITEM_CLICK时必须传入。 必须使用淘宝商品id，否则失败。
    itemId   string 

    // 游戏名称
    gameName   string 

    // 硬件CODE
    deviceCode   string 

    // 字段废弃
    endTime   string 

    // 字段废弃，考虑兼容，等同于op_time，两个必须传一个
    startTime   string 

    // ACTION枚举值：  BODY_SENSOR（通过人体感应、人脸识别成功识别到人） PHOTO_CLICK（用户在屏幕拍摄照片） GET_PHOTO（用户扫码获取照片，必须设置user_nick） ITEM_CLICK（商品点击时必须设置ITEM_ID）  GAME_START（开始游戏，可以 设置 user）  GAME_OVER_WITHOUT_PROMOTION（游戏结束）  WINNING_PRIZE（中奖，必须设置DRAW_RESULT） SHARE_CLICK（点击分享） RECEIVE_COUPONS (扫码领取优惠券时必须设置COUPON_ID)
    action   string 

    // 有则传入，没有可以不传。"例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
    couponId   string 

    // 用户昵称，混淆的淘宝NIck
    userNick   string 

    // 抽奖结果 ，如果传入，0：表示没中奖，1：表示中奖。该值必须是0或者1，传入其他失败。
    drawResult   string 

    // 数据外部编码，保证数据唯一性
    outerBizId   string 

    // 操作时间，后续统一使用该字段，考虑兼容，start_time跟该字段含义一致
    opTime   string 

    // 硬件识别的用户标识
    outerUser   string 

}

func NewTaobaoSmartstoreDeviceStoregameFeedbackRequest() *TaobaoSmartstoreDeviceStoregameFeedbackRequest{
    return &TaobaoSmartstoreDeviceStoregameFeedbackRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetApiMethodName() string {
    return "taobao.smartstore.device.storegame.feedback"
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSmartstoreDeviceStoregameFeedbackRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetItemId() string {
    return r.itemId
}

func (r *TaobaoSmartstoreDeviceStoregameFeedbackRequest) SetGameName(gameName string) error {
    r.gameName = gameName
    r.Set("game_name", gameName)
    return nil
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetGameName() string {
    return r.gameName
}

func (r *TaobaoSmartstoreDeviceStoregameFeedbackRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetDeviceCode() string {
    return r.deviceCode
}

func (r *TaobaoSmartstoreDeviceStoregameFeedbackRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoSmartstoreDeviceStoregameFeedbackRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSmartstoreDeviceStoregameFeedbackRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetAction() string {
    return r.action
}

func (r *TaobaoSmartstoreDeviceStoregameFeedbackRequest) SetCouponId(couponId string) error {
    r.couponId = couponId
    r.Set("coupon_id", couponId)
    return nil
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetCouponId() string {
    return r.couponId
}

func (r *TaobaoSmartstoreDeviceStoregameFeedbackRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetUserNick() string {
    return r.userNick
}

func (r *TaobaoSmartstoreDeviceStoregameFeedbackRequest) SetDrawResult(drawResult string) error {
    r.drawResult = drawResult
    r.Set("draw_result", drawResult)
    return nil
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetDrawResult() string {
    return r.drawResult
}

func (r *TaobaoSmartstoreDeviceStoregameFeedbackRequest) SetOuterBizId(outerBizId string) error {
    r.outerBizId = outerBizId
    r.Set("outer_biz_id", outerBizId)
    return nil
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetOuterBizId() string {
    return r.outerBizId
}

func (r *TaobaoSmartstoreDeviceStoregameFeedbackRequest) SetOpTime(opTime string) error {
    r.opTime = opTime
    r.Set("op_time", opTime)
    return nil
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetOpTime() string {
    return r.opTime
}

func (r *TaobaoSmartstoreDeviceStoregameFeedbackRequest) SetOuterUser(outerUser string) error {
    r.outerUser = outerUser
    r.Set("outer_user", outerUser)
    return nil
}

func (r TaobaoSmartstoreDeviceStoregameFeedbackRequest) GetOuterUser() string {
    return r.outerUser
}

