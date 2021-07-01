package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest
获取上游温度信息（疫苗） API请求
alibaba.alihealth.drug.kyt.dr.getupteminfo

根据追溯码及企业ID获取上游运输及存储温度信息（疫苗） */
type AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 企业ID
	_refEntId string
}

// New
