package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytUploadretailAPIRequest
门店销售扫码回传接口 API请求
alibaba.alihealth.drug.kyt.uploadretail

门店在销售给顾客时，扫描追溯码的数据按照单据回传； */
type AlibabaAlihealthDrugKytUploadretailAPIRequest struct {
	model.Params
	// 单据编号（唯一）
	_billCode string
	// 单据生成时间
	_billTime string
	// 单据类型[321,零售出库][322,疫苗接种]
	_billType int64
	// 药品类型[3,普药]
	_physicType int64
	// 码上放心平台企业编码（门店）
	_refUserId string
	// 发货企业(可为空)
	_fromUserId string
	// 单据提交者(appkey编号)
	_operIcCode string
	// 单据提交者姓名(可为空)
	_operIcName string
	// 请求类型[暂定都写2]
	_clientType string
	// 追溯码[多个时用逗号分开]
	_traceCodes []string
	// 患者电话
	_userTel string
	// 互联网标识
	_networkBillFlag string
	// 医师名单
	_medicDoctor string
	// 药品分发者
	_medicDispenser string
	// 患者名称
	_userName string
	// 患者代理领药人
	_userAgent string
	// 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
	_customerIdType string
	// 购买人证件编号
	_customerId string
}

// New
