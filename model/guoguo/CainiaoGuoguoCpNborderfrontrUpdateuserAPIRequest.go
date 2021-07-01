package guoguo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest
小件员信息变更 API请求
cainiao.guoguo.cp.nborderfrontr.updateuser

小件员信息变更 */
type CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest struct {
	model.Params
	// 姓名
	_name string
	// 网点站点信息
	_workStationName string
	// 小件员员工编号
	_cpUserId string
	// 支付宝账号
	_alipayAccount string
	// 城市
	_cityName string
	// 城市行政区域编码
	_cityCode string
	// 网点站点编码
	_workStationCode string
	// 小件员所在公司编号
	_cpCode string
	// 手机号
	_mobile string
}

// New
