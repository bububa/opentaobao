package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxDealsGetAPIRequest
批量获取交易列表 API请求
taobao.tanx.deals.get

批量获取交易信息 */
type TaobaoTanxDealsGetAPIRequest struct {
	model.Params
	// dsp用户id
	_dspId int64
	// dsp用户验证token
	_token string
	// 页大小
	_pageSize int64
	// 交易类型
	_dealType int64
	// 页码
	_page int64
	// 1970年到现在的时间，毫秒
	_signTime int64
}

// New
