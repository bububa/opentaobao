package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkActivityInfoGetAPIRequest
淘宝客-推广者-官方活动转链 API请求
taobao.tbk.activity.info.get

支持入参推广位和官方活动会场ID，获取活动信息和推广链接，包含推广长链接、短链接、淘口令、微信推广二维码地址等。改该API支持二方、三方类型的转链。官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取。 */
type TaobaoTbkActivityInfoGetAPIRequest struct {
	model.Params
	// mm_xxx_xxx_xxx的第三位
	_adzoneId int64
	// mm_xxx_xxx_xxx 仅三方分成场景使用
	_subPid string
	// 渠道关系id
	_relationId int64
	// 官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取
	_activityMaterialId string
	// 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
	_unionId string
}

// New
