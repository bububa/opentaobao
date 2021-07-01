package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxCreativeAddAPIRequest
创意预审新增接口 API请求
taobao.tanx.creative.add

创意预审新增接口 */
type TaobaoTanxCreativeAddAPIRequest struct {
	model.Params
	// DSP的memberId
	_memberId int64
	// dsp用户身份认证的TOKEN
	_token string
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
	// 创意id
	_creativeId string
	// 广告类目 如果有多个，以逗号分隔
	_adboardType string
	// 敏感词类目，多个的话以逗号分隔
	_sensitiveType string
	// 创意代码
	_adboardData string
	// 目标地址
	_destinationUrl string
	// 创意尺寸,中间以小写字母x分隔
	_adboardSize string
	// 创意封装类型：1 Htmlsnippet(pc网页),2 vast-nonlinear（视频暂停）,3vast-linear（视频贴片） 4 vast-companion(视频伴随),5 Html5(无线banner或app)
	_creativePackageFormat int64
}

// New
