package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytUpstorebillfileAPIRequest
上传零售出入库单(上传文件) API请求
alibaba.alihealth.drug.kyt.upstorebillfile

上传零售出入库单(上传文件) */
type AlibabaAlihealthDrugKytUpstorebillfileAPIRequest struct {
	model.Params
	// 单据编号
	_billCode string
	// 单据日期
	_billTime string
	// 单据类型[321,零售出库][322,疫苗接种]
	_billType int64
	// 药品类型[3,普药]
	_physicType int64
	// 上传企业的单位编码
	_refUserId string
	// 发货企业(参与人标识，为null时会通过refEntId自动得到)
	_fromUserId string
	// 单据提交者(key编号)
	_operIcCode string
	// 据提交者姓名
	_operIcName string
	// 文件内容
	_fileContent string
	// 文件名
	_uploadFileName string
	// 客户端类型[暂定都写2]
	_clientType string
}

// New
