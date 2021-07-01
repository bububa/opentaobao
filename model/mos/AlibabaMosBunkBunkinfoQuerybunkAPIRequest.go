package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosBunkBunkinfoQuerybunkAPIRequest
根据合同号查询铺位信息 API请求
alibaba.mos.bunk.bunkinfo.querybunk

根据合同号查询铺位信息 */
type AlibabaMosBunkBunkinfoQuerybunkAPIRequest struct {
	model.Params
	// 门店号
	_storeNo string
	// 合同状态集合
	_statusList []string
	// 合同号集合
	_contractCodes []string
}

// New
