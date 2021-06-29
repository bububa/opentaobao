package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-超级红包领取状态查询 API请求
taobao.tbk.dg.vegas.send.status

淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包（淘客接入前需签署协议 https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
*/
type TaobaoTbkDgVegasSendStatusRequest struct {
    model.Params
    // 渠道管理id
    relationId   string
    // 会员运营id
    specialId   string
    // 加密后的值(ALIPAY_ID除外)，需用MD5加密，32位小写
    deviceValue   string
    // 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE 5. ALIPAY_ID
    deviceType   string
    // thor平台业务码， 1：coupon 超红
    thorBizCode   string
    // 媒体pid
    pid   string
}

// 初始化TaobaoTbkDgVegasSendStatusRequest对象
func NewTaobaoTbkDgVegasSendStatusRequest() *TaobaoTbkDgVegasSendStatusRequest{
    return &TaobaoTbkDgVegasSendStatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgVegasSendStatusRequest) GetApiMethodName() string {
    return "taobao.tbk.dg.vegas.send.status"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgVegasSendStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RelationId Setter
// 渠道管理id
func (r *TaobaoTbkDgVegasSendStatusRequest) SetRelationId(relationId string) error {
    r.relationId = relationId
    r.Set("relation_id", relationId)
    return nil
}

// RelationId Getter
func (r TaobaoTbkDgVegasSendStatusRequest) GetRelationId() string {
    return r.relationId
}
// SpecialId Setter
// 会员运营id
func (r *TaobaoTbkDgVegasSendStatusRequest) SetSpecialId(specialId string) error {
    r.specialId = specialId
    r.Set("special_id", specialId)
    return nil
}

// SpecialId Getter
func (r TaobaoTbkDgVegasSendStatusRequest) GetSpecialId() string {
    return r.specialId
}
// DeviceValue Setter
// 加密后的值(ALIPAY_ID除外)，需用MD5加密，32位小写
func (r *TaobaoTbkDgVegasSendStatusRequest) SetDeviceValue(deviceValue string) error {
    r.deviceValue = deviceValue
    r.Set("device_value", deviceValue)
    return nil
}

// DeviceValue Getter
func (r TaobaoTbkDgVegasSendStatusRequest) GetDeviceValue() string {
    return r.deviceValue
}
// DeviceType Setter
// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE 5. ALIPAY_ID
func (r *TaobaoTbkDgVegasSendStatusRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r TaobaoTbkDgVegasSendStatusRequest) GetDeviceType() string {
    return r.deviceType
}
// ThorBizCode Setter
// thor平台业务码， 1：coupon 超红
func (r *TaobaoTbkDgVegasSendStatusRequest) SetThorBizCode(thorBizCode string) error {
    r.thorBizCode = thorBizCode
    r.Set("thor_biz_code", thorBizCode)
    return nil
}

// ThorBizCode Getter
func (r TaobaoTbkDgVegasSendStatusRequest) GetThorBizCode() string {
    return r.thorBizCode
}
// Pid Setter
// 媒体pid
func (r *TaobaoTbkDgVegasSendStatusRequest) SetPid(pid string) error {
    r.pid = pid
    r.Set("pid", pid)
    return nil
}

// Pid Getter
func (r TaobaoTbkDgVegasSendStatusRequest) GetPid() string {
    return r.pid
}
