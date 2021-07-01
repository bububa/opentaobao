package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstVendingCargospaceSaveAPIRequest
自动售卖机货道数据回流 API请求
alibaba.lst.vending.cargospace.save

自动售卖机货道数据回流接口，ISV通过调用此接口上传售卖机货道信息。 */
type AlibabaLstVendingCargospaceSaveAPIRequest struct {
	model.Params
	// 货道信息
	_cargoSpaceDTOList []VendingCargoSpaceDto
}

// New
