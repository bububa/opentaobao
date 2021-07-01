package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseContractdownloadAPIRequest
天猫开新车租后合同下载 API请求
tmall.car.lease.contractdownload

天猫开新车租后合同下载 */
type TmallCarLeaseContractdownloadAPIRequest struct {
	model.Params
	// 天猫开新车订单id
	_orderId int64
	// 续租协议： 1， 全款购车协议： 2，分期购买协议：3， 分期购买车辆资产验收协议：4，分期购买车辆抵押：5， 分期购买融资租赁合同：6
	_type string
}

// New
