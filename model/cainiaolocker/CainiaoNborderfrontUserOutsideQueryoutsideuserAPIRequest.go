package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest
查询外部小件员休息 API请求
cainiao.nborderfront.user.outside.queryoutsideuser

采用SPI方式查询外部公司的小件员信息 */
type CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest struct {
	model.Params
	// cpcode
	_cpCode string
	// cp小件员ID
	_cpUserId string
}

// New
