package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkerDeleteAPIRequest
删除工人 API请求
tmall.servicecenter.worker.delete

删除工人信息。该接口为多个业务公用，部分字段可忽略。对于电器预约安装业务，同一个服务商，通过工人姓名+手机号+biz_type 保证唯一性。工人已存在才可以删除。
错误码如下
100000, 系统错误
100001, 工人信息校验失败
100002, 用户校验失败
100003, 操作失败
10004,工人信息为空
10005,服务商id为空或者服务商名称为空
10006, 工人不存在
10007, 工人已存在
10008, 缺少工人姓名
10009, 缺少工人电话
10010, 网点不存在
11000, category_id 无效
11001, biz_type 无效
20001,已查询到最后一页 */
type TmallServicecenterWorkerDeleteAPIRequest struct {
	model.Params
	// 工人姓名
	_name string
	// 工人手机号
	_phone int64
	// 业务类型,电器预约安装业务填appliance_install
	_bizType string
}

// New
