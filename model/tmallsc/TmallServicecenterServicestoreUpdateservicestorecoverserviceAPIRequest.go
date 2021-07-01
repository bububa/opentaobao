package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest
更新网点覆盖的服务 API请求
tmall.servicecenter.servicestore.updateservicestorecoverservice

更新网点覆盖的服务，唯一性校验：服务商淘宝账号+网点编码+biz_type
前提是网点要存在，
如果需要新增的网点覆盖的服务不存在，会更新失败。
网点覆盖的服务包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、授权的类目和品牌 */
type TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest struct {
	model.Params
	// 业务类型
	_bizType string
	// json格式，在某个业务类型(biz_type)下类目和品牌的授权关系,比如空调授权了格力和美的，热水器授权了美的和林内，洗衣机和冰箱都授权了美的和松下
	_categoryIdsAndBrandIds string
	// serviceCodes列表,|分隔
	_serviceCodes string
	// 网点编码
	_serviceStoreCode string
}

// New
