package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest
药店出入库信息上传 API请求
alibaba.alihealth.drug.kyt.smyx.uploadinoutbill

药店上传出入库信息，包括采购入库（102），退货入库（103），供应入库（107）,退货出库（202），销毁出库（205），抽检出库（206）， 供应出库（209）,
不包括对个人的零售出库，疫苗接种，领药出库。 */
type AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest struct {
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
	// 发货企业entId
	_fromUserId string
	// 收货企业entId
	_toUserId string
	// 直调企业标识
	_destUserId string
	// 单据提交者（appkey编号）
	_operIcCode string
	// 单据提交者姓名
	_operIcName string
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
	// 仓号
	_warehouseId string
	// 药品ID[企业自已系统的药品ID]
	_drugId string
	// 追溯码[多个时用逗号分开]
	_traceCodes []string
}

// New
