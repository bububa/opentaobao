package elife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoElifeLifecardReconAPIRequest
查询对账文件地址接口 API请求
taobao.elife.lifecard.recon

查询对账文件地址接口 */
type TaobaoElifeLifecardReconAPIRequest struct {
	model.Params
	// 对账日期(YYYYMMDD)
	_opDate string
}

// New
