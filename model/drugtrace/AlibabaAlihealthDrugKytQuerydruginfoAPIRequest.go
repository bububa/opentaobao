package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytQuerydruginfoAPIRequest
码查询药品 API请求
alibaba.alihealth.drug.kyt.querydruginfo

通过追溯码查询药品信息 */
type AlibabaAlihealthDrugKytQuerydruginfoAPIRequest struct {
	model.Params
	// 码列表
	_codeList []string
	// 物流企业refentid
	_wuliuRefEntId string
	// 生产企业refentid
	_huozhuRefEntId string
}

// New
