package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsAddressReachablebatchGetAPIRequest
批量判定服务是否可达 API请求
taobao.logistics.address.reachablebatch.get

批量判定服务是否可达 */
type TaobaoLogisticsAddressReachablebatchGetAPIRequest struct {
	model.Params
	// 筛单用户输入地址结构
	_addressList []AddressReachable
}

// New
