package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscpublisherinfosaveAPIRequest 淘宝客-公用-私域用户备案 API请求
// taobao.tbk.sc.publisher.info.save
//
// 通过入参渠道管理或会员运营管理的邀请码，生成渠道id或会员运营id，完成渠道或会员的备案。
type TaobaotbkscpublisherinfosaveAPIRequest struct {
	model.Params
	// 渠道备案 - 来源，取链接的来源
	_relationFrom string
	// 渠道备案 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他
	_offlineScene string
	// 渠道备案 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他
	_onlineScene string
	// 淘宝客邀请渠道或会员的邀请码
	_inviterCode string
	// 媒体侧渠道备注
	_note string
	// 线下备案注册信息,字段包含: 电话号码(phoneNumber，必填),省(province,必填),市(city,必填),区县街道(location,必填),详细地址(detailAddress,必填),经营类型(career,线下个人必填),店铺类型(shopType,线下店铺必填),店铺名称(shopName,线下店铺必填),店铺证书类型(shopCertifyType,线下店铺选填),店铺证书编号(certifyNumber,线下店铺选填)
	_registerInfo string
	// 类型，必选 默认为1:
	_infoType int64
}

// NewTaobaotbkscpublisherinfosaveRequest 初始化TaobaotbkscpublisherinfosaveAPIRequest对象
func NewTaobaotbkscpublisherinfosaveRequest() *TaobaotbkscpublisherinfosaveAPIRequest {
	return &TaobaotbkscpublisherinfosaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.publisher.info.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationFrom is RelationFrom Setter
// 渠道备案 - 来源，取链接的来源
func (r *TaobaotbkscpublisherinfosaveAPIRequest) SetRelationFrom(_relationFrom string) error {
	r._relationFrom = _relationFrom
	r.Set("relation_from", _relationFrom)
	return nil
}

// GetRelationFrom RelationFrom Getter
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetRelationFrom() string {
	return r._relationFrom
}

// SetOfflineScene is OfflineScene Setter
// 渠道备案 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他
func (r *TaobaotbkscpublisherinfosaveAPIRequest) SetOfflineScene(_offlineScene string) error {
	r._offlineScene = _offlineScene
	r.Set("offline_scene", _offlineScene)
	return nil
}

// GetOfflineScene OfflineScene Getter
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetOfflineScene() string {
	return r._offlineScene
}

// SetOnlineScene is OnlineScene Setter
// 渠道备案 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他
func (r *TaobaotbkscpublisherinfosaveAPIRequest) SetOnlineScene(_onlineScene string) error {
	r._onlineScene = _onlineScene
	r.Set("online_scene", _onlineScene)
	return nil
}

// GetOnlineScene OnlineScene Getter
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetOnlineScene() string {
	return r._onlineScene
}

// SetInviterCode is InviterCode Setter
// 淘宝客邀请渠道或会员的邀请码
func (r *TaobaotbkscpublisherinfosaveAPIRequest) SetInviterCode(_inviterCode string) error {
	r._inviterCode = _inviterCode
	r.Set("inviter_code", _inviterCode)
	return nil
}

// GetInviterCode InviterCode Getter
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetInviterCode() string {
	return r._inviterCode
}

// SetNote is Note Setter
// 媒体侧渠道备注
func (r *TaobaotbkscpublisherinfosaveAPIRequest) SetNote(_note string) error {
	r._note = _note
	r.Set("note", _note)
	return nil
}

// GetNote Note Getter
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetNote() string {
	return r._note
}

// SetRegisterInfo is RegisterInfo Setter
// 线下备案注册信息,字段包含: 电话号码(phoneNumber，必填),省(province,必填),市(city,必填),区县街道(location,必填),详细地址(detailAddress,必填),经营类型(career,线下个人必填),店铺类型(shopType,线下店铺必填),店铺名称(shopName,线下店铺必填),店铺证书类型(shopCertifyType,线下店铺选填),店铺证书编号(certifyNumber,线下店铺选填)
func (r *TaobaotbkscpublisherinfosaveAPIRequest) SetRegisterInfo(_registerInfo string) error {
	r._registerInfo = _registerInfo
	r.Set("register_info", _registerInfo)
	return nil
}

// GetRegisterInfo RegisterInfo Getter
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetRegisterInfo() string {
	return r._registerInfo
}

// SetInfoType is InfoType Setter
// 类型，必选 默认为1:
func (r *TaobaotbkscpublisherinfosaveAPIRequest) SetInfoType(_infoType int64) error {
	r._infoType = _infoType
	r.Set("info_type", _infoType)
	return nil
}

// GetInfoType InfoType Getter
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetInfoType() int64 {
	return r._infoType
}
