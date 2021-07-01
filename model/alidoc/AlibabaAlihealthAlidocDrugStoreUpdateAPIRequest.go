package alidoc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest
更新药店 API请求
alibaba.alihealth.alidoc.drug.store.update

药店信息更新接口 */
type AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest struct {
	model.Params
	// 更新对象
	_drugStoreUpdateTopRequest *DrugStoreUpdateTopRequest
}

// New
