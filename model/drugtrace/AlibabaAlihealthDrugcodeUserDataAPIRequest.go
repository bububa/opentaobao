package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeUserDataAPIRequest
西安杨森同步用户行为接口 API请求
alibaba.alihealth.drugcode.user.data

西安杨森同步用户行为接口 */
type AlibabaAlihealthDrugcodeUserDataAPIRequest struct {
	model.Params
	// 用户信息
	_list []HaoxinqingDataDto
	// 企业ID，用户区分 appkey下不同企业数据隔离的
	_refEntId string
}

// New
