package ihome

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest
售后上门信息同步 API请求
alibaba.ihome.ctom.postsale.onsite.sync

用于三维家同步售后单上门人员和时间信息 */
type AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest struct {
	model.Params
	// 预约后展示给客户的员工联系方式
	_deliveryPhone string
	// 提交部门ID（预约的操作人所属部门）
	_unitId string
	// 备注
	_memo string
	// 预约事件发生时间
	_time int64
	// 服务ID
	_serviceId string
	// 预约目标时间
	_deliveryDate int64
	// 具体操作人ID（预约人）
	_operatorId string
	// 联系人员名字
	_deliveryName string
	// 售后单ID
	_postSalesId string
	// 配送、安装或上门
	_type string
	// 三维家补单ID
	_additionalOrderId string
}

// New
