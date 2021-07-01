package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationGoodsOnofflineAPIRequest
上线/下线 体检产品 API请求
alibaba.alihealth.examination.goods.onoffline

第三方体检机构对接钉钉体检中的产品 上线／下线 */
type AlibabaAlihealthExaminationGoodsOnofflineAPIRequest struct {
	model.Params
	// 商品组code，机构保证唯一
	_groupId string
	// 操作类型: online=上线，offline=下线
	_type string
	// 门店code列表
	_hospitalCodes string
}

// New
