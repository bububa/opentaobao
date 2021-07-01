package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxQualificationModifyAPIRequest
修改资质接口 API请求
taobao.tanx.qualification.modify

对dsp上传过的资质进行修改 */
type TaobaoTanxQualificationModifyAPIRequest struct {
	model.Params
	// dsp客户新增资质dto
	_qualifications []Qualification
	// dsp用户id
	_memberId int64
	// dsp用户验证token
	_token string
	// 1970年到现在的秒
	_signTime int64
}

// New
