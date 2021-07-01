package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugtraceTopLsydUploadretailAPIRequest
零售单据上传接口 API请求
alibaba.alihealth.drugtrace.top.lsyd.uploadretail

快易通多融零售上传接口 */
type AlibabaAlihealthDrugtraceTopLsydUploadretailAPIRequest struct {
	model.Params
	// 单据编号（唯一）
	_billCode string
	// 单据生成时间（一般写当前时间）
	_billTime string
	// 单据类型[321,零售出库][322,疫苗接种]
	_billType int64
	// 药品类型[2,特药，3,普药]
	_physicType int64
	// 码上放心平台企业唯一编码（门店或医疗机构）
	_refUserId string
	// 发货企业(可为空)
	_fromUserId string
	// 单据提交者(appkey编号)
	_operIcCode string
	// 单据提交者姓名(可为空)
	_operIcName string
	// 20位追溯码（多个时用半角逗号分隔）
	_traceCodes []string
	// 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
	_customerIdType string
	// 购买人证件编号
	_customerId string
	// 购买人电话
	_userTel string
	// 互联网订单标识 0非互联网 1互联网
	_networkBillFlag string
	// 开药医师
	_medicDoctor string
	// 药品发药人
	_medicDispenser string
	// 药品使用者姓名
	_userName string
	// 药品代理人
	_userAgent string
}

// New
