package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxQualificationSolidFindAPIRequest
客户固态共享资质 API请求
taobao.tanx.qualification.solid.find

接口会返回该广告主下的所有审核通过并且可被共享的资质，这些资质在过期之前可以不需要再次上传。 */
type TaobaoTanxQualificationSolidFindAPIRequest struct {
	model.Params
	// 广告主id
	_advertiserId int64
	// 资质元素id列表
	_elementIds []int64
	// dsp用户id
	_memberId int64
	// dsp客户验证token
	_token string
	// 1970年到现在的秒
	_signTime int64
	// 查询起始页
	_page int64
	// 分页大小
	_pageSize int64
}

// New
