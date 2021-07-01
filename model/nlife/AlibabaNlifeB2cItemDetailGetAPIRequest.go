package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2cItemDetailGetAPIRequest
b2c码详情查询 API请求
alibaba.nlife.b2c.item.detail.get

根据零售+平台生成的唯一码获取对应详情 */
type AlibabaNlifeB2cItemDetailGetAPIRequest struct {
	model.Params
	// 商家入驻门店在零售+平台的ID
	_storeId string
	// 零售+平台生成的唯一码或条码
	_uniqueCode string
}

// New
