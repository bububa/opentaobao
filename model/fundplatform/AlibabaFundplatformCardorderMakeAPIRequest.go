package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformCardorderMakeAPIRequest
通知制卡商制卡 API请求
alibaba.fundplatform.cardorder.make

该接口由内部定义，外部制卡商实现。当需要制卡商进行制卡操作时，将会调用该接口。 */
type AlibabaFundplatformCardorderMakeAPIRequest struct {
	model.Params
	// 物流信息
	_logisticsInfo *AlibabaFundplatformCardorderMakeStruct
	// 卡模板信息列表
	_cardProductInfos []AlibabaFundplatformCardorderMakeStruct
	// 子制卡单ID
	_cardOrderId int64
	// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
	_ownSign string
}

// New
