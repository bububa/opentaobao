package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScPublisherInfoSaveAPIRequest 淘宝客-公用-私域用户备案 API请求
// taobao.tbk.sc.publisher.info.save
//
// 通过入参渠道管理或会员运营管理的邀请码，生成渠道id或会员运营id，完成渠道或会员的备案。
type TaobaoTbkScPublisherInfoSaveAPIRequest struct {
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

// NewTaobaoTbkScPublisherInfoSaveRequest 初始化TaobaoTbkScPublisherInfoSaveAPIRequest对象
func NewTaobaoTbkScPublisherInfoSaveRequest() *TaobaoTbkScPublisherInfoSaveAPIRequest {
	return &TaobaoTbkScPublisherInfoSaveAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkScPublisherInfoSaveAPIRequest) Reset() {
	r._relationFrom = ""
	r._offlineScene = ""
	r._onlineScene = ""
	r._inviterCode = ""
	r._note = ""
	r._registerInfo = ""
	r._infoType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScPublisherInfoSaveAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.publisher.info.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScPublisherInfoSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScPublisherInfoSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationFrom is RelationFrom Setter
// 渠道备案 - 来源，取链接的来源
func (r *TaobaoTbkScPublisherInfoSaveAPIRequest) SetRelationFrom(_relationFrom string) error {
	r._relationFrom = _relationFrom
	r.Set("relation_from", _relationFrom)
	return nil
}

// GetRelationFrom RelationFrom Getter
func (r TaobaoTbkScPublisherInfoSaveAPIRequest) GetRelationFrom() string {
	return r._relationFrom
}

// SetOfflineScene is OfflineScene Setter
// 渠道备案 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他
func (r *TaobaoTbkScPublisherInfoSaveAPIRequest) SetOfflineScene(_offlineScene string) error {
	r._offlineScene = _offlineScene
	r.Set("offline_scene", _offlineScene)
	return nil
}

// GetOfflineScene OfflineScene Getter
func (r TaobaoTbkScPublisherInfoSaveAPIRequest) GetOfflineScene() string {
	return r._offlineScene
}

// SetOnlineScene is OnlineScene Setter
// 渠道备案 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他
func (r *TaobaoTbkScPublisherInfoSaveAPIRequest) SetOnlineScene(_onlineScene string) error {
	r._onlineScene = _onlineScene
	r.Set("online_scene", _onlineScene)
	return nil
}

// GetOnlineScene OnlineScene Getter
func (r TaobaoTbkScPublisherInfoSaveAPIRequest) GetOnlineScene() string {
	return r._onlineScene
}

// SetInviterCode is InviterCode Setter
// 淘宝客邀请渠道或会员的邀请码
func (r *TaobaoTbkScPublisherInfoSaveAPIRequest) SetInviterCode(_inviterCode string) error {
	r._inviterCode = _inviterCode
	r.Set("inviter_code", _inviterCode)
	return nil
}

// GetInviterCode InviterCode Getter
func (r TaobaoTbkScPublisherInfoSaveAPIRequest) GetInviterCode() string {
	return r._inviterCode
}

// SetNote is Note Setter
// 媒体侧渠道备注
func (r *TaobaoTbkScPublisherInfoSaveAPIRequest) SetNote(_note string) error {
	r._note = _note
	r.Set("note", _note)
	return nil
}

// GetNote Note Getter
func (r TaobaoTbkScPublisherInfoSaveAPIRequest) GetNote() string {
	return r._note
}

// SetRegisterInfo is RegisterInfo Setter
// 线下备案注册信息,字段包含: 电话号码(phoneNumber，必填),省(province,必填),市(city,必填),区县街道(location,必填),详细地址(detailAddress,必填),经营类型(career,线下个人必填),店铺类型(shopType,线下店铺必填),店铺名称(shopName,线下店铺必填),店铺证书类型(shopCertifyType,线下店铺选填),店铺证书编号(certifyNumber,线下店铺选填)
func (r *TaobaoTbkScPublisherInfoSaveAPIRequest) SetRegisterInfo(_registerInfo string) error {
	r._registerInfo = _registerInfo
	r.Set("register_info", _registerInfo)
	return nil
}

// GetRegisterInfo RegisterInfo Getter
func (r TaobaoTbkScPublisherInfoSaveAPIRequest) GetRegisterInfo() string {
	return r._registerInfo
}

// SetInfoType is InfoType Setter
// 类型，必选 默认为1:
func (r *TaobaoTbkScPublisherInfoSaveAPIRequest) SetInfoType(_infoType int64) error {
	r._infoType = _infoType
	r.Set("info_type", _infoType)
	return nil
}

// GetInfoType InfoType Getter
func (r TaobaoTbkScPublisherInfoSaveAPIRequest) GetInfoType() int64 {
	return r._infoType
}

var poolTaobaoTbkScPublisherInfoSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkScPublisherInfoSaveRequest()
	},
}

// GetTaobaoTbkScPublisherInfoSaveRequest 从 sync.Pool 获取 TaobaoTbkScPublisherInfoSaveAPIRequest
func GetTaobaoTbkScPublisherInfoSaveAPIRequest() *TaobaoTbkScPublisherInfoSaveAPIRequest {
	return poolTaobaoTbkScPublisherInfoSaveAPIRequest.Get().(*TaobaoTbkScPublisherInfoSaveAPIRequest)
}

// ReleaseTaobaoTbkScPublisherInfoSaveAPIRequest 将 TaobaoTbkScPublisherInfoSaveAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkScPublisherInfoSaveAPIRequest(v *TaobaoTbkScPublisherInfoSaveAPIRequest) {
	v.Reset()
	poolTaobaoTbkScPublisherInfoSaveAPIRequest.Put(v)
}
