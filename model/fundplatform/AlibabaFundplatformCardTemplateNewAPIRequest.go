package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformCardTemplateNewAPIRequest
新增实体卡模板 API请求
alibaba.fundplatform.card.template.new

该接口由制卡商实现，当新增一个实体卡模板的时候，需要调用该接口，通知制卡商同步新增卡模板信息。 */
type AlibabaFundplatformCardTemplateNewAPIRequest struct {
	model.Params
	// 卡模板编号
	_templateNo string
	// 该模板生成的卡名称
	_cardName string
	// 卡面额，单元分
	_parValue string
	// 卡外观图片地址
	_pictureUrl string
	// 是否为测试卡模板，true表示是，如果是测试卡模板则请求制卡时无需真正去制作实体卡
	_isTest bool
	// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
	_ownSign string
}

// New
