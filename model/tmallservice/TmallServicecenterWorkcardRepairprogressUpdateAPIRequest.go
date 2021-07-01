package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardRepairprogressUpdateAPIRequest
更新维修进度 API请求
tmall.servicecenter.workcard.repairprogress.update

提供给外部合作服务商的维修进度更改接口 */
type TmallServicecenterWorkcardRepairprogressUpdateAPIRequest struct {
	model.Params
	// 图片列表
	_picUrlList []string
	// 请求节点的动作描述，唯一标识一个节点
	_action string
	// 工单id
	_workcardId int64
	// 真实接单服务商账号Nick
	_realTpNick string
	// 服务目标物瑕疵信息
	_targetGoodsDefects string
	// 衣服，鞋子
	_receivedGoods string
}

// New
