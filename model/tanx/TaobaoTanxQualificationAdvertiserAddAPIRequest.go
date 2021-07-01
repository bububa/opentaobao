package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxQualificationAdvertiserAddAPIRequest
新增广告主接口 API请求
taobao.tanx.qualification.advertiser.add

外部dsp调用接口时会根据广告主名称和类型在tanx系统中新增一个广告主 */
type TaobaoTanxQualificationAdvertiserAddAPIRequest struct {
	model.Params
	// 广告主对象
	_advertisers []AdvertiserDto
	// dsp用户memberId
	_memberId int64
	// dsp用户验证token
	_token string
	// 从1970年到当前时间的秒
	_signTime int64
}

// New
