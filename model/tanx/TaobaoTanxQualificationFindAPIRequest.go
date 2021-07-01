package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxQualificationFindAPIRequest
资质查询接口 API请求
taobao.tanx.qualification.find

资质查询接口 */
type TaobaoTanxQualificationFindAPIRequest struct {
	model.Params
	// dsp客户在tanx的memberId
	_memberId int64
	// dsp客户在tanx签名的token
	_token string
	// 当前时间,从1970年算的毫秒
	_signTime int64
	// 分页的第几页，从1开始
	_page int64
	// 分页大小，限制200
	_pageSize int64
	// 广告主资质查询dto
	_query *QualificationQuery
}

// New
