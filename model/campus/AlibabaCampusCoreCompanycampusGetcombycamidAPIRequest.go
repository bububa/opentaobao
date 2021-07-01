package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest
根据园区ID获取运营公司信息 API请求
alibaba.campus.core.companycampus.getcombycamid

根据园区ID获取运营公司信息 */
type AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest struct {
	model.Params
	// WorkBenchContext
	_param0 *WorkBenchContext
}

// New
