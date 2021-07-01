package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardStatusUpdateAPIRequest
服务商反馈服务的执行情况 API请求
tmall.servicecenter.workcard.status.update

1 如果服务商受理了此服务，修改合同状态为：已受理=3<br/>2 如果服务商没有受理此服务，修改合同状态为：已拒绝=10<br/>3 如果服务商执行了此服务，修改合同状态为：已执行=4<br/>4 如果服务商执行服务成功，修改合同状态为：已完成=5<br/>5 如果此工单为合同类型的工单，当服务商受理了此服务后，会进行分账 */
type TmallServicecenterWorkcardStatusUpdateAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 工单类型： 2（合同） 或者 1(任务）
	_type *model.File
	// 目前仅支持5种状态的反馈：3=授理， 10=拒绝 ，4=执行 ，5=成功，11=失败。（所有状态列表： -1： 初始化 0： 生成 1： 生效 2： 申请 3： 受理 4： 执行 5： 成功 9： 结算 10： 拒绝 11： 失败 12 ： 撤销 13： 暂停 19： 终止）
	_status *model.File
	// 买家id
	_buyerId int64
	// api调用者
	_updater string
	// 更新时间
	_updateDate int64
	// 服务生效时间 ：工单类型为合同工单时，必选！
	_effectDate int64
	// 服务失效时间 ：工单类型为合同工单时，必选！
	_expireDate int64
	// 备注,256个字符以内
	_comments string
	// 任务类工单，预约或者上门地址
	_address string
	// 任务执行，预约联系人
	_contactName string
	// 任务执行，预约联系人电话
	_contactPhone string
	// 服务预约时间
	_serviceDate int64
	// 服务完成时间
	_completeDate int64
	// 服务凭证上传的图片URL链接，多个以;隔开
	_serviceVoucherPics string
	// 属性定义。例如无忧退货服务，K-V对定义，每对KV用“;”分割，“:”号左边是key右边是value，value如果有多个则以“,”分割。 reasons   :  原因，可能有多个 succeedCount     :    取件成功个数 failedCount    :    取件失败个数 cancelCount      :     取件取消个数 totalCount       :      总取件个数，totalCount= succeedCount + failedCount + cancelCount
	_attribute string
	// 服务商网点内部编码
	_serviceCenterCode string
	// 服务商网点名字
	_serviceCenterName string
	// 单元是分
	_serviceFee int64
	// 是否上门
	_isVisit bool
	// 说明
	_beforeServiceMemo string
	// 说明
	_afterServiceMemo string
	// 手机号码
	_phoneImei string
	// 服务子状态：30 表示“服务已申请（上门）” 31表示“服务改约（上门）” 400表示“服务结果（待件上门）” 410表示“服务结果（拖机维修）” 411表示“服务结果（换机）” 420表示“服务结果（上门不可维修）”
	_subStatus int64
	// 网点负责人联系电话
	_serviceCenterManagerPhone string
	// 网点负责人
	_serviceCenterManagerName string
	// 网点地址
	_serviceCenterAddress string
	// 一个工单可能包含多件商品，比如空调可能有两台，录入每天机器的安装情况
	_workCardInstallDetailList []WorkCardInstallDetail
	// json string。费用单位为分
	_serviceFeeDetail string
	// 物流单号
	_expressCode string
	// 物流公司名字
	_expressCompany string
}

// New
