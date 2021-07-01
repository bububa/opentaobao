package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkScPublisherInfoSaveAPIRequest
淘宝客-公用-私域用户备案 API请求
taobao.tbk.sc.publisher.info.save

通过入参渠道管理或会员运营管理的邀请码，生成渠道id或会员运营id，完成渠道或会员的备案。 */
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
	// 类型，必选 默认为1:
	_infoType int64
	// 媒体侧渠道备注
	_note string
	// 线下备案注册信息,字段包含: 电话号码(phoneNumber，必填),省(province,必填),市(city,必填),区县街道(location,必填),详细地址(detailAddress,必填),经营类型(career,线下个人必填),店铺类型(shopType,线下店铺必填),店铺名称(shopName,线下店铺必填),店铺证书类型(shopCertifyType,线下店铺选填),店铺证书编号(certifyNumber,线下店铺选填)
	_registerInfo string
}

// New
