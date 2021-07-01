package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytUpinoutfileAPIRequest
上传出入库单据(传文件) API请求
alibaba.alihealth.drug.kyt.upinoutfile

上传出入库单据(传文件) */
type AlibabaAlihealthDrugKytUpinoutfileAPIRequest struct {
	model.Params
	// 单据编码
	_billCode string
	// 单据时间
	_billTime string
	// 单据类型【102代表采购入库】
	_billType int64
	// 药品类型【3普药2特药】
	_physicType int64
	// 上传企业的单位编码
	_refUserId string
	// 代理企业REF标识
	_agentRefUserId string
	// 收货企业entId
	_fromUserId string
	// 发货企业entId
	_toUserId string
	// 直调企业标识
	_destUserId string
	// 单据提交者（key编号）
	_operIcCode string
	// 单据提交者姓名
	_operIcName string
	// 仓号
	_warehouseId string
	// 药品ID
	_drugId string
	// 文件内容
	_fileContent string
	// 文件名
	_uploadFileName string
	// 客户端类型[必须填2]
	_clientType string
	// 退货原因代码[退货入出库时填写]
	_returnReasonCode string
	// 退货原因描述[退货入出库时填写]
	_returnReasonDes string
	// 注销原因代码【销毁出库时填写】
	_cancelReasonCode string
	// 注销原因描述【销毁出库时填写】
	_cancelReasonDes string
	// 执行人姓名【销毁出库时填写】
	_executerName string
	// 执行人证件号【销毁出库时填写】
	_executerCode string
	// 监督人姓名【销毁出库时填写】
	_superviserName string
	// 监督人证件号【销毁出库时填写】
	_superviserCode string
}

// New
