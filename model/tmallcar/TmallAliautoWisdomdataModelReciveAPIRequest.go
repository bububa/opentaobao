package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallAliautoWisdomdataModelReciveAPIRequest
第三方车型数据上传 API请求
tmall.aliauto.wisdomdata.model.recive

天猫汽车对外提供的汽车车型数据上传接口 */
type TmallAliautoWisdomdataModelReciveAPIRequest struct {
	model.Params
	// JSON格式车型完整数据
	_modelDetail string
	// 接入的第三方库中的车型唯一id
	_resourceId string
	// 接入的第三方库的名称
	_fromSource string
}

// New
