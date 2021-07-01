package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxCreativesGetAPIRequest
批量获取DSP用户的创意审核结果 API请求
taobao.tanx.creatives.get

批量获取DSP用户的创意审核结果 */
type TaobaoTanxCreativesGetAPIRequest struct {
	model.Params
	// DSP的memberId
	_memberId int64
	// dsp用户身份认证的TOKEN
	_token string
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
	// 创意的状态（全部ALL,通过PASS,拒绝REFUSE,未审核WAITING）
	_status string
	// 分页的页码(第一页为1)
	_page int64
	// 所选创意的类型。1-->普通类型, 2-->视频贴片, 0 -->优先查询普通类型,无结果则查询视频贴片类型
	_type int64
}

// New
