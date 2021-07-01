package dt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDtTmllcarPricevalidateAPIRequest
线索报价价格校验 API请求
alibaba.dt.tmllcar.pricevalidate

根据选定的车型和城市，校验汽车价格是否通过
入参：车型ID，城市名称，价格
输出：N 校验失败，校验成功不返回值 */
type AlibabaDtTmllcarPricevalidateAPIRequest struct {
	model.Params
	// tt
	_appName string
	// tt
	_name string
	// tt
	_password string
	// tt
	_price *BigDecimal
	// tt
	_cityName string
	// tt
	_modelName string
}

// New
