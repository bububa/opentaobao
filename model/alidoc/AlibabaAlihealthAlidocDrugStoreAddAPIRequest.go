package alidoc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAlidocDrugStoreAddAPIRequest
gsk新增药店 API请求
alibaba.alihealth.alidoc.drug.store.add

GSK上传药店信息 */
type AlibabaAlihealthAlidocDrugStoreAddAPIRequest struct {
	model.Params
	// 新增药店
	_drugStoreAddTopRequest *DrugStoreAddTopRequest
}

// New
