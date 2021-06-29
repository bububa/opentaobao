package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-私域用户备案 API请求
taobao.tbk.sc.publisher.info.save

通过入参渠道管理或会员运营管理的邀请码，生成渠道id或会员运营id，完成渠道或会员的备案。
*/
type TaobaoTbkScPublisherInfoSaveRequest struct {
    model.Params
    // 渠道备案 - 来源，取链接的来源
    _relationFrom   string
    // 渠道备案 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他
    _offlineScene   string
    // 渠道备案 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他
    _onlineScene   string
    // 淘宝客邀请渠道或会员的邀请码
    _inviterCode   string
    // 类型，必选 默认为1:
    _infoType   int64
    // 媒体侧渠道备注
    _note   string
    // 线下备案注册信息,字段包含: 电话号码(phoneNumber，必填),省(province,必填),市(city,必填),区县街道(location,必填),详细地址(detailAddress,必填),经营类型(career,线下个人必填),店铺类型(shopType,线下店铺必填),店铺名称(shopName,线下店铺必填),店铺证书类型(shopCertifyType,线下店铺选填),店铺证书编号(certifyNumber,线下店铺选填)
    _registerInfo   string
}

// 初始化TaobaoTbkScPublisherInfoSaveRequest对象
func NewTaobaoTbkScPublisherInfoSaveRequest() *TaobaoTbkScPublisherInfoSaveRequest{
    return &TaobaoTbkScPublisherInfoSaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkScPublisherInfoSaveRequest) GetApiMethodName() string {
    return "taobao.tbk.sc.publisher.info.save"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkScPublisherInfoSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RelationFrom Setter
// 渠道备案 - 来源，取链接的来源
func (r *TaobaoTbkScPublisherInfoSaveRequest) SetRelationFrom(_relationFrom string) error {
    r._relationFrom = _relationFrom
    r.Set("relation_from", _relationFrom)
    return nil
}

// RelationFrom Getter
func (r TaobaoTbkScPublisherInfoSaveRequest) GetRelationFrom() string {
    return r._relationFrom
}
// OfflineScene Setter
// 渠道备案 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他
func (r *TaobaoTbkScPublisherInfoSaveRequest) SetOfflineScene(_offlineScene string) error {
    r._offlineScene = _offlineScene
    r.Set("offline_scene", _offlineScene)
    return nil
}

// OfflineScene Getter
func (r TaobaoTbkScPublisherInfoSaveRequest) GetOfflineScene() string {
    return r._offlineScene
}
// OnlineScene Setter
// 渠道备案 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他
func (r *TaobaoTbkScPublisherInfoSaveRequest) SetOnlineScene(_onlineScene string) error {
    r._onlineScene = _onlineScene
    r.Set("online_scene", _onlineScene)
    return nil
}

// OnlineScene Getter
func (r TaobaoTbkScPublisherInfoSaveRequest) GetOnlineScene() string {
    return r._onlineScene
}
// InviterCode Setter
// 淘宝客邀请渠道或会员的邀请码
func (r *TaobaoTbkScPublisherInfoSaveRequest) SetInviterCode(_inviterCode string) error {
    r._inviterCode = _inviterCode
    r.Set("inviter_code", _inviterCode)
    return nil
}

// InviterCode Getter
func (r TaobaoTbkScPublisherInfoSaveRequest) GetInviterCode() string {
    return r._inviterCode
}
// InfoType Setter
// 类型，必选 默认为1:
func (r *TaobaoTbkScPublisherInfoSaveRequest) SetInfoType(_infoType int64) error {
    r._infoType = _infoType
    r.Set("info_type", _infoType)
    return nil
}

// InfoType Getter
func (r TaobaoTbkScPublisherInfoSaveRequest) GetInfoType() int64 {
    return r._infoType
}
// Note Setter
// 媒体侧渠道备注
func (r *TaobaoTbkScPublisherInfoSaveRequest) SetNote(_note string) error {
    r._note = _note
    r.Set("note", _note)
    return nil
}

// Note Getter
func (r TaobaoTbkScPublisherInfoSaveRequest) GetNote() string {
    return r._note
}
// RegisterInfo Setter
// 线下备案注册信息,字段包含: 电话号码(phoneNumber，必填),省(province,必填),市(city,必填),区县街道(location,必填),详细地址(detailAddress,必填),经营类型(career,线下个人必填),店铺类型(shopType,线下店铺必填),店铺名称(shopName,线下店铺必填),店铺证书类型(shopCertifyType,线下店铺选填),店铺证书编号(certifyNumber,线下店铺选填)
func (r *TaobaoTbkScPublisherInfoSaveRequest) SetRegisterInfo(_registerInfo string) error {
    r._registerInfo = _registerInfo
    r.Set("register_info", _registerInfo)
    return nil
}

// RegisterInfo Getter
func (r TaobaoTbkScPublisherInfoSaveRequest) GetRegisterInfo() string {
    return r._registerInfo
}
