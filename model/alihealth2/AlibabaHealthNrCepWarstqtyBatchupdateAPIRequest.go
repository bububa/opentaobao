package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest
批量更新ISV库存 API请求
alibaba.health.nr.cep.warstqty.batchupdate

青岛医保服务-ISV批量更新孔雀翎中库存数据 */
type AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest struct {
	model.Params
	// 库存更新对象
	_warStqtyList []TopIsvStqtyLstDto
}

// New
