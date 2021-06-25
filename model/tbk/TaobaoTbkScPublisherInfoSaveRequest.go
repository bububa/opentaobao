package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-私域用户备案 APIRequest
taobao.tbk.sc.publisher.info.save

通过入参渠道管理或会员运营管理的邀请码，生成渠道id或会员运营id，完成渠道或会员的备案。
*/
type TaobaoTbkScPublisherInfoSaveRequest struct {
    model.Params

    // 渠道备案 - 来源，取链接的来源
    relationFrom   string 

    // 渠道备案 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他
    offlineScene   string 

    // 渠道备案 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他
    onlineScene   string 

    // 淘宝客邀请渠道或会员的邀请码
    inviterCode   string 

    // 类型，必选 默认为1:
    infoType   int64 

    // 媒体侧渠道备注
    note   string 

    // 线下备案注册信息,字段包含: 电话号码(phoneNumber，必填),省(province,必填),市(city,必填),区县街道(location,必填),详细地址(detailAddress,必填),经营类型(career,线下个人必填),店铺类型(shopType,线下店铺必填),店铺名称(shopName,线下店铺必填),店铺证书类型(shopCertifyType,线下店铺选填),店铺证书编号(certifyNumber,线下店铺选填)
    registerInfo   string 

}

func NewTaobaoTbkScPublisherInfoSaveRequest() *TaobaoTbkScPublisherInfoSaveRequest{
    return &TaobaoTbkScPublisherInfoSaveRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkScPublisherInfoSaveRequest) GetApiMethodName() string {
    return "taobao.tbk.sc.publisher.info.save"
}

func (r TaobaoTbkScPublisherInfoSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkScPublisherInfoSaveRequest) SetRelationFrom(relationFrom string) error {
    r.relationFrom = relationFrom
    r.Set("relation_from", relationFrom)
    return nil
}

func (r TaobaoTbkScPublisherInfoSaveRequest) GetRelationFrom() string {
    return r.relationFrom
}

func (r *TaobaoTbkScPublisherInfoSaveRequest) SetOfflineScene(offlineScene string) error {
    r.offlineScene = offlineScene
    r.Set("offline_scene", offlineScene)
    return nil
}

func (r TaobaoTbkScPublisherInfoSaveRequest) GetOfflineScene() string {
    return r.offlineScene
}

func (r *TaobaoTbkScPublisherInfoSaveRequest) SetOnlineScene(onlineScene string) error {
    r.onlineScene = onlineScene
    r.Set("online_scene", onlineScene)
    return nil
}

func (r TaobaoTbkScPublisherInfoSaveRequest) GetOnlineScene() string {
    return r.onlineScene
}

func (r *TaobaoTbkScPublisherInfoSaveRequest) SetInviterCode(inviterCode string) error {
    r.inviterCode = inviterCode
    r.Set("inviter_code", inviterCode)
    return nil
}

func (r TaobaoTbkScPublisherInfoSaveRequest) GetInviterCode() string {
    return r.inviterCode
}

func (r *TaobaoTbkScPublisherInfoSaveRequest) SetInfoType(infoType int64) error {
    r.infoType = infoType
    r.Set("info_type", infoType)
    return nil
}

func (r TaobaoTbkScPublisherInfoSaveRequest) GetInfoType() int64 {
    return r.infoType
}

func (r *TaobaoTbkScPublisherInfoSaveRequest) SetNote(note string) error {
    r.note = note
    r.Set("note", note)
    return nil
}

func (r TaobaoTbkScPublisherInfoSaveRequest) GetNote() string {
    return r.note
}

func (r *TaobaoTbkScPublisherInfoSaveRequest) SetRegisterInfo(registerInfo string) error {
    r.registerInfo = registerInfo
    r.Set("register_info", registerInfo)
    return nil
}

func (r TaobaoTbkScPublisherInfoSaveRequest) GetRegisterInfo() string {
    return r.registerInfo
}

