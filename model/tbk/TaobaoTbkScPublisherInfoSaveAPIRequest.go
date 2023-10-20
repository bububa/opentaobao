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
	_relationfrom string
	// 渠道备案 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他
	_offlinescene string
	// 渠道备案 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他
	_onlinescene string
	// 淘宝客邀请渠道或会员的邀请码
	_invitercode string
	// 媒体侧渠道备注
	_note string
	// 线下备案注册信息,字段包含: 电话号码(phoneNumber，必填),省(province,必填),市(city,必填),区县街道(location,必填),详细地址(detailAddress,必填),经营类型(career,线下个人必填),店铺类型(shopType,线下店铺必填),店铺名称(shopName,线下店铺必填),店铺证书类型(shopCertifyType,线下店铺选填),店铺证书编号(certifyNumber,线下店铺选填)
	_registerinfo string
	// 类型，必选 默认为1:
	_infotype int64
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

// SetRelationfrom is Relationfrom Setter
// 渠道备案 - 来源，取链接的来源
func (r *TaobaotbkscpublisherinfosaveAPIRequest) SetRelationfrom(_relationfrom string) error {
	r._relationfrom = _relationfrom
	r.Set("relation_from", _relationfrom)
	return nil
}

// GetRelationfrom Relationfrom Getter
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetRelationfrom() string {
	return r._relationfrom
}

// SetOfflinescene is Offlinescene Setter
// 渠道备案 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他
func (r *TaobaotbkscpublisherinfosaveAPIRequest) SetOfflinescene(_offlinescene string) error {
	r._offlinescene = _offlinescene
	r.Set("offline_scene", _offlinescene)
	return nil
}

// GetOfflinescene Offlinescene Getter
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetOfflinescene() string {
	return r._offlinescene
}

// SetOnlinescene is Onlinescene Setter
// 渠道备案 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他
func (r *TaobaotbkscpublisherinfosaveAPIRequest) SetOnlinescene(_onlinescene string) error {
	r._onlinescene = _onlinescene
	r.Set("online_scene", _onlinescene)
	return nil
}

// GetOnlinescene Onlinescene Getter
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetOnlinescene() string {
	return r._onlinescene
}

// SetInvitercode is Invitercode Setter
// 淘宝客邀请渠道或会员的邀请码
func (r *TaobaotbkscpublisherinfosaveAPIRequest) SetInvitercode(_invitercode string) error {
	r._invitercode = _invitercode
	r.Set("inviter_code", _invitercode)
	return nil
}

// GetInvitercode Invitercode Getter
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetInvitercode() string {
	return r._invitercode
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

// SetRegisterinfo is Registerinfo Setter
// 线下备案注册信息,字段包含: 电话号码(phoneNumber，必填),省(province,必填),市(city,必填),区县街道(location,必填),详细地址(detailAddress,必填),经营类型(career,线下个人必填),店铺类型(shopType,线下店铺必填),店铺名称(shopName,线下店铺必填),店铺证书类型(shopCertifyType,线下店铺选填),店铺证书编号(certifyNumber,线下店铺选填)
func (r *TaobaotbkscpublisherinfosaveAPIRequest) SetRegisterinfo(_registerinfo string) error {
	r._registerinfo = _registerinfo
	r.Set("register_info", _registerinfo)
	return nil
}

// GetRegisterinfo Registerinfo Getter
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetRegisterinfo() string {
	return r._registerinfo
}

// SetInfotype is Infotype Setter
// 类型，必选 默认为1:
func (r *TaobaotbkscpublisherinfosaveAPIRequest) SetInfotype(_infotype int64) error {
	r._infotype = _infotype
	r.Set("info_type", _infotype)
	return nil
}

// GetInfotype Infotype Getter
func (r TaobaotbkscpublisherinfosaveAPIRequest) GetInfotype() int64 {
	return r._infotype
}
