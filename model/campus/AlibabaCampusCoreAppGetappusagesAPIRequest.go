package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusCoreAppGetappusagesAPIRequest
根据应用ID获得应用所在的园区 API请求
alibaba.campus.core.app.getappusages

传入应用的id,  获得用户授权的园区 */
type AlibabaCampusCoreAppGetappusagesAPIRequest struct {
	model.Params
	// 应用id
	_appid int64
	// WorkBenchContext
	_workBenchContext *WorkBenchContext
}

// New
