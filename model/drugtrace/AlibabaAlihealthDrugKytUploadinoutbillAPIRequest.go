package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytUploadinoutbillAPIRequest
企业上传出入库信息 API请求
alibaba.alihealth.drug.kyt.uploadinoutbill

企业上传出入库信息，包括101, "生产入库"；102, "采购入库"；103, "退货入库"；104, "调拨入库"；106, "零头入库"；107, "供应入库"；108, "召回入库"；110,"赠品入库"；111,"盘盈入库"；112,"报废入库"；113,"其他入库"
201, "销售出库"；202, "退货出库"；203, "调拨出库"；204, "返工出库"；205, "销毁出库"；206, "抽检出库"；207, "直调出库"；208, "生产出库"；209, "供应出库"；211, "召回出库"；212,"赠品出库"；214,"盘亏出库"；215,"损坏出库"；216,"报废出库"；217,"其他出库"；237, "直调退货"。
不包括对个人的零售出库，疫苗接种，领药出库。
本接口与uploadcircubill接口的主要区别的，本接口入参中直接上传追溯码（多个码时用逗号分隔）。uploadcircubill接口入参中，需要上传码的单据文件（用扫码枪生成的xml文件），一般情况下使用uploadcircubill接口上传单据文件。 */
type AlibabaAlihealthDrugKytUploadinoutbillAPIRequest struct {
	model.Params
	// 单据编号【同一个企业不能上传相同单据号】
	_billCode string
	// 单据时间（扫码时间）
	_billTime string
	// 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
	_billType int64
	// 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药
	_physicType int64
	// 货主（单据的所有者，上传人），上传企业的单位维一编码【注意：该入参是ref_ent_id，不是ent_id】
	_refUserId string
	// 第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
	_agentRefUserId string
	// 发货企业entId【注意：该入参是ent_id,并不是ref_ent_id】
	_fromUserId string
	// 收货企业entId【注意：该入参是ent_id,并不是ref_ent_id】
	_toUserId string
	// 直调企业标识【注意：该入参是ent_id,并不是ref_ent_id】
	_destUserId string
	// 单据提交者（调用接口时的appkey编号）
	_operIcCode string
	// 单据提交者姓名（出入库单上传人的名子）
	_operIcName string
	// 调用方类型[必须填2]
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
	// （协同平台数据合规）应收货总数量
	_quReceivable int64
	// （协同平台数据合规）是否验证，0：未通过验证，1：已验证
	_xtIsCheck string
	// （协同平台数据合规）未验证通过原因【验证未通过时填写】
	_xtCheckCode string
	// （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
	_xtCheckCodeDesc string
	// （协同平台数据合规）药品列表Json："codeCount": 药品数量 "commDrugId": 国家药品唯一标识 "exprieDate": 生产日期 "physicInfo": 药品信息 "pkgSpec": 包状规格 "prepnCount": 制剂数量 "produceBatchNo":生产批次 "produceDate": 生产日期
	_drugListJson string
	// （协同平台数据合规）单据委托企业refEntId
	_assRefEntId string
	// （协同平台数据合规）药品配送企业entId出库单，收货方为医疗机构时填写】
	_disEntId string
	// （协同平台数据合规）药品配送企业出库单，收货方为医疗机构时填写】
	_disRefEntId string
	// （协同平台数据合规）收货人【必选】
	_toPerson string
	// （协同平台数据合规）发货人【必选】
	_fromPerson string
	// （协同平台数据合规）订货单编号【可选】
	_orderCode string
	// （协同平台数据合规）发货单编号【必选】
	_fromBillCode string
	// （协同平台数据合规）收货地址【必选】
	_toAddress string
	// （协同平台数据合规）发货地址【必选】
	_fromAddress string
	// （协同平台数据合规）单据委托企业entId
	_assEntId string
}

// New
