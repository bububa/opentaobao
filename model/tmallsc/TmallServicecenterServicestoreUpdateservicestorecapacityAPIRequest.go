package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterServicestoreUpdateservicestorecapacityAPIRequest
更新网点容量 API请求
tmall.servicecenter.servicestore.updateservicestorecapacity

更新网点容量，唯一性校验：服务商淘宝账号+网点编码+biz_type
前提是网点要存在，
如果需要更新的网点容量不存在，会更新失败。
网点容量包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、类目区域和容量 */
type TmallServicecenterServicestoreUpdateservicestorecapacityAPIRequest struct {
	model.Params
	// 业务类型
	_bizType string
	// json格式，在某个业务类型(biz_type)下,类目所覆盖区域对应的容量。一个网点承接了空调和热水器的安装, 空调覆盖的区域是杭州市上城区和下城区，容量是10； 热水器覆盖的区域是杭州市下城区和江干区，容量是18；洗衣机和冰箱覆盖区域一样都是杭州市上城区和西湖区，合并计算容量30
	_categoryIdsAndAreaCodesAndCapacity string
	// serviceCodes列表,|分隔
	_serviceCodes string
	// 网点编码
	_serviceStoreCode string
}

// New
